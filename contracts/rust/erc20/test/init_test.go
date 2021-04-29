package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	erc20file = "erc20_bg.wasm"
)

func TestDeployErc20(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")

	creator, creatorAddr = env.NewKeyPairWithFunds()
	creatorAgentID = coretypes.NewAgentID(creatorAddr, 0)
	err := chain.DeployWasmContract(nil, ScName, erc20file,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.NoError(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))

	_, err = chain.FindContract(ScName)
	require.NoError(t, err)

	// deploy second time
	err = chain.DeployWasmContract(nil, ScName, erc20file,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec = chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))
}

func TestDeployErc20Fail1(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, ScName, erc20file)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail2(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, ScName, erc20file,
		ParamSupply, 1000000,
	)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail3(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	creator, creatorAddr = env.NewKeyPairWithFunds()
	creatorAgentID = coretypes.NewAgentID(creatorAddr, 0)
	err := chain.DeployWasmContract(nil, ScName, erc20file,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail3Repeat(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	creator, creatorAddr = env.NewKeyPairWithFunds()
	creatorAgentID = coretypes.NewAgentID(creatorAddr, 0)
	err := chain.DeployWasmContract(nil, ScName, erc20file,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))

	// repeat after failure
	err = chain.DeployWasmContract(nil, ScName, erc20file,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.NoError(t, err)
	_, _, rec = chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))

	_, err = chain.FindContract(ScName)
	require.NoError(t, err)
}
