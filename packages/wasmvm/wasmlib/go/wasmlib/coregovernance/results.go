// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package coregovernance

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

type ArrayOfImmutableAddress struct {
	Proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableAddress) Length() uint32 {
	return a.Proxy.Length()
}

func (a ArrayOfImmutableAddress) GetAddress(index uint32) wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(a.Proxy.Index(index))
}

type ImmutableGetAllowedStateControllerAddressesResults struct {
	Proxy wasmtypes.Proxy
}

// Array16 of state controller addresses
func (s ImmutableGetAllowedStateControllerAddressesResults) Controllers() ArrayOfImmutableAddress {
	return ArrayOfImmutableAddress{Proxy: s.Proxy.Root(ResultControllers)}
}

type ArrayOfMutableAddress struct {
	Proxy wasmtypes.Proxy
}

func (a ArrayOfMutableAddress) AppendAddress() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(a.Proxy.Append())
}

func (a ArrayOfMutableAddress) Clear() {
	a.Proxy.ClearArray()
}

func (a ArrayOfMutableAddress) Length() uint32 {
	return a.Proxy.Length()
}

func (a ArrayOfMutableAddress) GetAddress(index uint32) wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(a.Proxy.Index(index))
}

type MutableGetAllowedStateControllerAddressesResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetAllowedStateControllerAddressesResults() MutableGetAllowedStateControllerAddressesResults {
	return MutableGetAllowedStateControllerAddressesResults{Proxy: wasmlib.NewResultsProxy()}
}

// Array16 of state controller addresses
func (s MutableGetAllowedStateControllerAddressesResults) Controllers() ArrayOfMutableAddress {
	return ArrayOfMutableAddress{Proxy: s.Proxy.Root(ResultControllers)}
}

type ImmutableGetChainInfoResults struct {
	Proxy wasmtypes.Proxy
}

// chain ID
func (s ImmutableGetChainInfoResults) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.Proxy.Root(ResultChainID))
}

// chain owner agent ID
func (s ImmutableGetChainInfoResults) ChainOwnerID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ResultChainOwnerID))
}

// serialized fee policy
func (s ImmutableGetChainInfoResults) FeePolicy() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultFeePolicy))
}

// serialized gas limits
func (s ImmutableGetChainInfoResults) GasLimits() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultGasLimits))
}

// chain metadata
func (s ImmutableGetChainInfoResults) Metadata() ImmutableChainMetadata {
	return ImmutableChainMetadata{Proxy: s.Proxy.Root(ResultMetadata)}
}

func (s ImmutableGetChainInfoResults) PublicURL() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.Proxy.Root(ResultPublicURL))
}

type MutableGetChainInfoResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetChainInfoResults() MutableGetChainInfoResults {
	return MutableGetChainInfoResults{Proxy: wasmlib.NewResultsProxy()}
}

// chain ID
func (s MutableGetChainInfoResults) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.Proxy.Root(ResultChainID))
}

// chain owner agent ID
func (s MutableGetChainInfoResults) ChainOwnerID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ResultChainOwnerID))
}

// serialized fee policy
func (s MutableGetChainInfoResults) FeePolicy() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultFeePolicy))
}

// serialized gas limits
func (s MutableGetChainInfoResults) GasLimits() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultGasLimits))
}

// chain metadata
func (s MutableGetChainInfoResults) Metadata() MutableChainMetadata {
	return MutableChainMetadata{Proxy: s.Proxy.Root(ResultMetadata)}
}

func (s MutableGetChainInfoResults) PublicURL() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.Proxy.Root(ResultPublicURL))
}

type MapBytesToImmutableBytes struct {
	Proxy wasmtypes.Proxy
}

func (m MapBytesToImmutableBytes) GetBytes(key []byte) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.Proxy.Key(wasmtypes.BytesToBytes(key)))
}

type MapBytesToImmutableBool struct {
	Proxy wasmtypes.Proxy
}

func (m MapBytesToImmutableBool) GetBool(key []byte) wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(m.Proxy.Key(wasmtypes.BytesToBytes(key)))
}

type ImmutableGetChainNodesResults struct {
	Proxy wasmtypes.Proxy
}

// serialized access node info per pubKey
func (s ImmutableGetChainNodesResults) AccessNodeCandidates() MapBytesToImmutableBytes {
	return MapBytesToImmutableBytes{Proxy: s.Proxy.Root(ResultAccessNodeCandidates)}
}

// pubKey set
func (s ImmutableGetChainNodesResults) AccessNodes() MapBytesToImmutableBool {
	return MapBytesToImmutableBool{Proxy: s.Proxy.Root(ResultAccessNodes)}
}

type MapBytesToMutableBytes struct {
	Proxy wasmtypes.Proxy
}

func (m MapBytesToMutableBytes) Clear() {
	m.Proxy.ClearMap()
}

func (m MapBytesToMutableBytes) GetBytes(key []byte) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.Proxy.Key(wasmtypes.BytesToBytes(key)))
}

type MapBytesToMutableBool struct {
	Proxy wasmtypes.Proxy
}

func (m MapBytesToMutableBool) Clear() {
	m.Proxy.ClearMap()
}

func (m MapBytesToMutableBool) GetBool(key []byte) wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(m.Proxy.Key(wasmtypes.BytesToBytes(key)))
}

type MutableGetChainNodesResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetChainNodesResults() MutableGetChainNodesResults {
	return MutableGetChainNodesResults{Proxy: wasmlib.NewResultsProxy()}
}

// serialized access node info per pubKey
func (s MutableGetChainNodesResults) AccessNodeCandidates() MapBytesToMutableBytes {
	return MapBytesToMutableBytes{Proxy: s.Proxy.Root(ResultAccessNodeCandidates)}
}

// pubKey set
func (s MutableGetChainNodesResults) AccessNodes() MapBytesToMutableBool {
	return MapBytesToMutableBool{Proxy: s.Proxy.Root(ResultAccessNodes)}
}

type ImmutableGetChainOwnerResults struct {
	Proxy wasmtypes.Proxy
}

// chain owner
func (s ImmutableGetChainOwnerResults) ChainOwner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ResultChainOwner))
}

type MutableGetChainOwnerResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetChainOwnerResults() MutableGetChainOwnerResults {
	return MutableGetChainOwnerResults{Proxy: wasmlib.NewResultsProxy()}
}

// chain owner
func (s MutableGetChainOwnerResults) ChainOwner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ResultChainOwner))
}

type ImmutableGetEVMGasRatioResults struct {
	Proxy wasmtypes.Proxy
}

// serialized gas ratio
func (s ImmutableGetEVMGasRatioResults) GasRatio() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultGasRatio))
}

type MutableGetEVMGasRatioResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetEVMGasRatioResults() MutableGetEVMGasRatioResults {
	return MutableGetEVMGasRatioResults{Proxy: wasmlib.NewResultsProxy()}
}

// serialized gas ratio
func (s MutableGetEVMGasRatioResults) GasRatio() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultGasRatio))
}

type ImmutableGetFeePolicyResults struct {
	Proxy wasmtypes.Proxy
}

// serialized fee policy
func (s ImmutableGetFeePolicyResults) FeePolicy() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultFeePolicy))
}

type MutableGetFeePolicyResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetFeePolicyResults() MutableGetFeePolicyResults {
	return MutableGetFeePolicyResults{Proxy: wasmlib.NewResultsProxy()}
}

// serialized fee policy
func (s MutableGetFeePolicyResults) FeePolicy() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultFeePolicy))
}

type ImmutableGetGasLimitsResults struct {
	Proxy wasmtypes.Proxy
}

// serialized gas limits
func (s ImmutableGetGasLimitsResults) GasLimits() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultGasLimits))
}

type MutableGetGasLimitsResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetGasLimitsResults() MutableGetGasLimitsResults {
	return MutableGetGasLimitsResults{Proxy: wasmlib.NewResultsProxy()}
}

// serialized gas limits
func (s MutableGetGasLimitsResults) GasLimits() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultGasLimits))
}

type ImmutableGetMaintenanceStatusResults struct {
	Proxy wasmtypes.Proxy
}

// whether maintenance mode is on
func (s ImmutableGetMaintenanceStatusResults) Status() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.Proxy.Root(ResultStatus))
}

type MutableGetMaintenanceStatusResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetMaintenanceStatusResults() MutableGetMaintenanceStatusResults {
	return MutableGetMaintenanceStatusResults{Proxy: wasmlib.NewResultsProxy()}
}

// whether maintenance mode is on
func (s MutableGetMaintenanceStatusResults) Status() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.Proxy.Root(ResultStatus))
}

type ImmutableGetMetadataResults struct {
	Proxy wasmtypes.Proxy
}

// the L2 metadata
func (s ImmutableGetMetadataResults) Metadata() ImmutableChainMetadata {
	return ImmutableChainMetadata{Proxy: s.Proxy.Root(ResultMetadata)}
}

// the public url leading to the chain info, stored on the tangle (l1)
func (s ImmutableGetMetadataResults) PublicURL() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.Proxy.Root(ResultPublicURL))
}

type MutableGetMetadataResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetMetadataResults() MutableGetMetadataResults {
	return MutableGetMetadataResults{Proxy: wasmlib.NewResultsProxy()}
}

// the L2 metadata
func (s MutableGetMetadataResults) Metadata() MutableChainMetadata {
	return MutableChainMetadata{Proxy: s.Proxy.Root(ResultMetadata)}
}

// the public url leading to the chain info, stored on the tangle (l1)
func (s MutableGetMetadataResults) PublicURL() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.Proxy.Root(ResultPublicURL))
}
