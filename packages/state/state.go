// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package state

import (
	"fmt"
	"github.com/iotaledger/wasp/packages/kv/trie"
	"github.com/iotaledger/wasp/packages/kv/trie_merkle"
	"math"
	"time"

	"github.com/iotaledger/hive.go/kvstore"
	"github.com/iotaledger/hive.go/kvstore/mapdb"
	"github.com/iotaledger/wasp/packages/database/dbkeys"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/coreutil"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/buffered"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/util"
	"golang.org/x/xerrors"
)

// region VirtualStateAccess /////////////////////////////////////////////////

type virtualStateAccess struct {
	db   kvstore.KVStore
	kvs  *buffered.BufferedKVStoreAccess
	trie *trie.Trie
}

var (
	CommitmentModel                    = trie_merkle.Model
	_               VirtualStateAccess = &virtualStateAccess{}
)

// newVirtualState creates VirtualStateAccess interface with the partition of KVStore
func newVirtualState(db kvstore.KVStore) *virtualStateAccess {
	subState := subRealm(db, []byte{dbkeys.ObjectTypeState})
	subTrie := subRealm(db, []byte{dbkeys.ObjectTypeTrie})
	ret := &virtualStateAccess{
		db:   db,
		kvs:  buffered.NewBufferedKVStoreAccess(kv.NewHiveKVStoreReader(subState)),
		trie: trie.New(CommitmentModel, kv.NewHiveKVStoreReader(subTrie)),
	}
	return ret
}

// CreateOriginState origin state and saves it. It assumes store is empty
func newOriginState(store kvstore.KVStore) VirtualStateAccess {
	ret := newVirtualState(store)
	nilChainId := iscp.ChainID{}
	// state will contain chain ID at key ''. In the origin state it 'all 0'
	var zero8 [8]byte
	ret.KVStore().Set("", nilChainId.Bytes())
	ret.KVStore().Set(kv.Key(coreutil.StatePrefixBlockIndex), zero8[:])
	ret.KVStore().Set(kv.Key(coreutil.StatePrefixTimestamp), zero8[:])
	ret.Commit()
	return ret
}

// calcOriginStateHash is independent of db provider nor chainID. Used for testing
func calcOriginStateHash() trie.VCommitment {
	return trie.RootCommitment(newOriginState(mapdb.NewMapDB()).TrieAccess())
}

// CreateOriginState creates and saves origin state in DB
func CreateOriginState(store kvstore.KVStore, chainID *iscp.ChainID) (VirtualStateAccess, error) {
	originState := newOriginState(store)
	if err := originState.Save(); err != nil {
		return nil, err
	}
	// state will contain chain ID at key ''.
	// We set the mutation, but we do not commit yet, it will be committed with the block #1
	originState.KVStore().Set("", chainID.Bytes())
	return originState, nil
}

func subRealm(db kvstore.KVStore, realm []byte) kvstore.KVStore {
	if db == nil {
		return nil
	}
	return db.WithRealm(append(db.Realm(), realm...))
}

func (vs *virtualStateAccess) Copy() VirtualStateAccess {
	ret := &virtualStateAccess{
		db:   vs.db,
		kvs:  vs.kvs.Copy(),
		trie: vs.trie.Clone(),
	}
	return ret
}

func (vs *virtualStateAccess) DangerouslyConvertToString() string {
	return fmt.Sprintf("#%d, ts: %v, state commitment: %s\n%s",
		vs.BlockIndex(),
		vs.Timestamp(),
		trie.RootCommitment(vs.TrieAccess()),
		vs.KVStore().DangerouslyDumpToString(),
	)
}

func (vs *virtualStateAccess) KVStore() *buffered.BufferedKVStoreAccess {
	return vs.kvs
}

func (vs *virtualStateAccess) KVStoreReader() kv.KVStoreReader {
	return vs.kvs
}

func (vs *virtualStateAccess) BlockIndex() uint32 {
	blockIndex, err := loadStateIndexFromState(vs.kvs)
	if err != nil {
		panic(xerrors.Errorf("state.BlockIndex: %w", err))
	}
	return blockIndex
}

func (vs *virtualStateAccess) Timestamp() time.Time {
	ts, err := loadTimestampFromState(vs.kvs)
	if err != nil {
		panic(xerrors.Errorf("state.OutputTimestamp: %w", err))
	}
	return ts
}

func (vs *virtualStateAccess) PreviousStateCommitment() trie.VCommitment {
	cBin, err := vs.KVStore().Get(kv.Key(coreutil.StatePrefixPrevStateCommitment))
	if err != nil {
		panic(xerrors.Errorf("state.PreviousStateCommitment: %w", err))
	}
	c, err := vs.trie.VectorCommitmentFromBytes(cBin)
	if err != nil {
		panic(xerrors.Errorf("loadPrevStateHashFromState: %w", err))
	}
	return c
}

// ApplyBlock applies a block of state updates. Checks consistency of the block and previous state. Updates state hash
// It is not suitible for applying origin block to empty virtual state. This is done in `newZeroVirtualState`
func (vs *virtualStateAccess) ApplyBlock(b Block) error {
	if vs.BlockIndex()+1 != b.BlockIndex() {
		return xerrors.Errorf("ApplyBlock: b state index #%d can't be applied to the state with index #%d",
			b.BlockIndex(), vs.BlockIndex())
	}
	if vs.Timestamp().After(b.Timestamp()) {
		return xerrors.New("ApplyBlock: inconsistent timestamps")
	}
	vs.applyBlockNoCheck(b)
	return nil
}

func (vs *virtualStateAccess) applyBlockNoCheck(b Block) {
	vs.ApplyStateUpdate(b.(*blockImpl).stateUpdate)
}

// ApplyStateUpdate applies one state update
func (vs *virtualStateAccess) ApplyStateUpdate(upd Update) {
	upd.Mutations().Apply(vs.kvs)
}

func (vs *virtualStateAccess) ProofGeneric(key []byte) *trie.ProofGeneric {
	return trie.GetProofGeneric(vs.trie, dbkeys.MakeKey(dbkeys.ObjectTypeTrie, key))
}

// ExtractBlock creates a block from update log and returns it or nil if log is empty. The log is cleared
func (vs *virtualStateAccess) ExtractBlock() (Block, error) {
	ret, err := newBlock(vs.kvs.Mutations())
	if err != nil {
		return nil, err
	}
	if vs.BlockIndex() != ret.BlockIndex() {
		return nil, xerrors.New("virtualStateAccess: internal inconsistency: index of the state is not equal to the index of the extracted block")
	}
	return ret, nil
}

func (vs *virtualStateAccess) Commit() {
	if !vs.kvs.Mutations().IsModified() {
		return
	}
	for k, v := range vs.kvs.Mutations().Sets {
		vs.trie.Update([]byte(k), v)
	}
	for k := range vs.kvs.Mutations().Dels {
		vs.trie.Update([]byte(k), nil)
	}
	vs.trie.Commit()
	vs.kvs.Mutations().ResetModified()
}

// StateCommitment returns the hash of the state, calculated as a hashing of the previous (committed) state hash and the block hash.
func (vs *virtualStateAccess) TrieAccess() trie.NodeStore {
	return vs.trie
}

// ReconcileTrie a heavy operation
func (vs *virtualStateAccess) ReconcileTrie() []kv.Key {
	return vs.trie.Reconcile(vs.kvs)
}

func loadStateIndexFromState(chainState kv.KVStoreReader) (uint32, error) {
	blockIndexBin, err := chainState.Get(kv.Key(coreutil.StatePrefixBlockIndex))
	if err != nil {
		return 0, err
	}
	if blockIndexBin == nil {
		return 0, xerrors.New("loadStateIndexFromState: not found")
	}
	blockIndex, err := util.Uint64From8Bytes(blockIndexBin)
	if err != nil {
		return 0, xerrors.Errorf("loadStateIndexFromState: %w", err)
	}
	if int(blockIndex) > math.MaxUint32 {
		return 0, xerrors.Errorf("loadStateIndexFromState: wrong state index value")
	}
	return uint32(blockIndex), nil
}

func loadTimestampFromState(chainState kv.KVStoreReader) (time.Time, error) {
	tsBin, err := chainState.Get(kv.Key(coreutil.StatePrefixTimestamp))
	if err != nil {
		return time.Time{}, err
	}
	ts, err := codec.DecodeTime(tsBin)
	if err != nil {
		return time.Time{}, xerrors.Errorf("loadTimestampFromState: %w", err)
	}
	return ts, nil
}
