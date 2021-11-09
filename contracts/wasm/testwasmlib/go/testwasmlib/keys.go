// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamAddress     = 0
	IdxParamAgentID     = 1
	IdxParamBlockIndex  = 2
	IdxParamBytes       = 3
	IdxParamChainID     = 4
	IdxParamColor       = 5
	IdxParamHash        = 6
	IdxParamHname       = 7
	IdxParamIndex       = 8
	IdxParamInt16       = 9
	IdxParamInt32       = 10
	IdxParamInt64       = 11
	IdxParamName        = 12
	IdxParamParam       = 13
	IdxParamRecordIndex = 14
	IdxParamRequestID   = 15
	IdxParamString      = 16
	IdxParamValue       = 17
	IdxResultCount      = 18
	IdxResultIotas      = 19
	IdxResultLength     = 20
	IdxResultRandom     = 21
	IdxResultRecord     = 22
	IdxResultValue      = 23
	IdxStateArrays      = 24
	IdxStateRandom      = 25
)

const keyMapLen = 26

var keyMap = [keyMapLen]wasmlib.Key{
	ParamAddress,
	ParamAgentID,
	ParamBlockIndex,
	ParamBytes,
	ParamChainID,
	ParamColor,
	ParamHash,
	ParamHname,
	ParamIndex,
	ParamInt16,
	ParamInt32,
	ParamInt64,
	ParamName,
	ParamParam,
	ParamRecordIndex,
	ParamRequestID,
	ParamString,
	ParamValue,
	ResultCount,
	ResultIotas,
	ResultLength,
	ResultRandom,
	ResultRecord,
	ResultValue,
	StateArrays,
	StateRandom,
}

var idxMap [keyMapLen]wasmlib.Key32