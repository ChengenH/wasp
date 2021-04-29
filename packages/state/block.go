package state

import (
	"bytes"
	"fmt"
	"io"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/kvstore"
	"github.com/iotaledger/wasp/packages/dbprovider"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/util"
)

type block struct {
	stateIndex    uint32
	stateOutputID ledgerstate.OutputID
	stateUpdates  []StateUpdate
}

const OriginBlockHashBase58 = "6dHpJfMrjZsMtLasN2zbQNfzxMvoP2hNbLA95Xv2JHAN"

// validates, enumerates and creates a block from array of state updates
func NewBlock(stateUpdates ...StateUpdate) (Block, error) {
	if len(stateUpdates) == 0 {
		return nil, fmt.Errorf("block can't be empty")
	}
	return &block{
		stateUpdates: stateUpdates,
	}, nil
}

func BlockFromBytes(data []byte) (Block, error) {
	ret := new(block)
	if err := ret.Read(bytes.NewReader(data)); err != nil {
		return nil, err
	}
	return ret, nil
}

func LoadBlock(partition kvstore.KVStore, stateIndex uint32) (Block, error) {
	data, err := partition.Get(dbkeyBlock(stateIndex))
	if err == kvstore.ErrKeyNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return BlockFromBytes(data)
}

// block with empty state update and nil state hash
func MustNewOriginBlock() Block {
	ret, err := NewBlock(NewStateUpdate())
	if err != nil {
		log.Panic(err)
	}
	return ret
}

func (b *block) Bytes() []byte {
	var buf bytes.Buffer
	_ = b.Write(&buf)
	return buf.Bytes()
}

func (b *block) String() string {
	ret := ""
	ret += fmt.Sprintf("Block: state index: %d\n", b.StateIndex())
	ret += fmt.Sprintf("state txid: %s\n", b.ApprovingOutputID().String())
	ret += fmt.Sprintf("timestamp: %d\n", b.Timestamp())
	ret += fmt.Sprintf("size: %d\n", b.Size())
	ret += fmt.Sprintf("essence: %s\n", b.EssenceHash().String())
	for i, su := range b.stateUpdates {
		ret += fmt.Sprintf("   #%d: %s\n", i, su.String())
	}
	return ret
}

func (b *block) ApprovingOutputID() ledgerstate.OutputID {
	return b.stateOutputID
}

func (b *block) StateIndex() uint32 {
	return b.stateIndex
}

// timestmap of the last state update
func (b *block) Timestamp() int64 {
	return b.stateUpdates[len(b.stateUpdates)-1].Timestamp()
}

func (b *block) WithBlockIndex(stateIndex uint32) Block {
	b.stateIndex = stateIndex
	return b
}

func (b *block) WithApprovingOutputID(vtxid ledgerstate.OutputID) Block {
	b.stateOutputID = vtxid
	return b
}

func (b *block) ForEach(fun func(uint16, StateUpdate) bool) {
	for i, su := range b.stateUpdates {
		if !fun(uint16(i), su) {
			return
		}
	}
}

func (b *block) Size() uint16 {
	return uint16(len(b.stateUpdates))
}

// hash of all data except state transaction hash
func (b *block) EssenceHash() hashing.HashValue {
	var buf bytes.Buffer
	if err := b.writeEssence(&buf); err != nil {
		panic("EssenceHash")
	}
	return hashing.HashData(buf.Bytes())
}

func (b *block) Write(w io.Writer) error {
	if err := b.writeEssence(w); err != nil {
		return err
	}
	if _, err := w.Write(b.stateOutputID.Bytes()); err != nil {
		return err
	}
	return nil
}

func (b *block) writeEssence(w io.Writer) error {
	if err := util.WriteUint32(w, b.stateIndex); err != nil {
		return err
	}
	if err := util.WriteUint16(w, uint16(len(b.stateUpdates))); err != nil {
		return err
	}
	for _, su := range b.stateUpdates {
		if err := su.Write(w); err != nil {
			return err
		}
	}
	return nil
}

func (b *block) Read(r io.Reader) error {
	if err := b.readEssence(r); err != nil {
		return err
	}
	if _, err := r.Read(b.stateOutputID[:]); err != nil {
		return err
	}
	return nil
}

func (b *block) readEssence(r io.Reader) error {
	if err := util.ReadUint32(r, &b.stateIndex); err != nil {
		return err
	}
	var size uint16
	if err := util.ReadUint16(r, &size); err != nil {
		return err
	}
	b.stateUpdates = make([]StateUpdate, size)
	var err error
	for i := range b.stateUpdates {
		b.stateUpdates[i], err = NewStateUpdateRead(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func dbkeyBlock(stateIndex uint32) []byte {
	return dbprovider.MakeKey(dbprovider.ObjectTypeStateUpdateBatch, util.Uint32To4Bytes(stateIndex))
}

func (b *block) IsApprovedBy(chainOutput *ledgerstate.AliasOutput) bool {
	if chainOutput == nil {
		return false
	}
	if b.StateIndex() != chainOutput.GetStateIndex() {
		return false
	}
	var nilOID ledgerstate.OutputID
	if b.ApprovingOutputID() != nilOID && b.ApprovingOutputID() != chainOutput.ID() {
		return false
	}
	sh, err := hashing.HashValueFromBytes(chainOutput.GetStateData())
	if err != nil {
		return false
	}
	if b.EssenceHash() != sh {
		return false
	}
	return true
}
