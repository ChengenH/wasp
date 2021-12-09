package utxodb

import (
	"testing"
	"time"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/iota.go/v3/tpkg"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	u := New()
	require.EqualValues(t, u.Supply(), u.GetAddressBalanceIotas(u.GenesisAddress()))
	gtime := u.GlobalTime()
	require.EqualValues(t, 1, gtime.MilestoneIndex)
	expectedTime := time.Unix(1, 0).Add(1 * time.Millisecond)
	require.EqualValues(t, expectedTime, gtime.Time)

	u.AdvanceClockBy(10*time.Second, 3)
	gtime1 := u.GlobalTime()
	require.EqualValues(t, 4, gtime1.MilestoneIndex)
	expectedTime = gtime.Time.Add(10 * time.Second)
	require.EqualValues(t, expectedTime, gtime1.Time)
}

func TestRequestFunds(t *testing.T) {
	u := New()
	addr := tpkg.RandEd25519Address()
	tx, err := u.RequestFunds(addr)
	require.NoError(t, err)
	require.EqualValues(t, u.Supply()-RequestFundsAmount, u.GetAddressBalanceIotas(u.GenesisAddress()))
	require.EqualValues(t, RequestFundsAmount, u.GetAddressBalanceIotas(addr))

	txID, err := tx.ID()
	require.NoError(t, err)
	require.Same(t, tx, u.MustGetTransaction(*txID))

	gtime := u.GlobalTime()
	require.EqualValues(t, 2, gtime.MilestoneIndex)
	expectedTime := time.Unix(1, 0).Add(2 * time.Millisecond)
	require.EqualValues(t, expectedTime, gtime.Time)
}

func TestAddTransactionFail(t *testing.T) {
	u := New()

	addr := tpkg.RandEd25519Address()
	tx, err := u.RequestFunds(addr)
	require.NoError(t, err)

	err = u.AddTransaction(tx)
	require.Error(t, err)
}

func TestDoubleSpend(t *testing.T) {
	keyPair1 := cryptolib.NewKeyPair()

	addr1 := cryptolib.Ed25519AddressFromPubKey(keyPair1.PublicKey)
	key1Signer := iotago.NewInMemoryAddressSigner(iotago.NewAddressKeysForEd25519Address(addr1, keyPair1.PrivateKey))

	addr2 := tpkg.RandEd25519Address()
	addr3 := tpkg.RandEd25519Address()

	u := New()

	tx1, err := u.RequestFunds(addr1)
	require.NoError(t, err)
	tx1ID, err := tx1.ID()
	require.NoError(t, err)

	spend2, err := iotago.NewTransactionBuilder().
		AddInput(&iotago.ToBeSignedUTXOInput{Address: addr1, Input: &iotago.UTXOInput{
			TransactionID:          *tx1ID,
			TransactionOutputIndex: 0,
		}}).
		AddOutput(&iotago.ExtendedOutput{Address: addr2, Amount: RequestFundsAmount}).
		Build(u.deSeriParams(), key1Signer)
	require.NoError(t, err)
	err = u.AddTransaction(spend2)
	require.NoError(t, err)

	spend3, err := iotago.NewTransactionBuilder().
		AddInput(&iotago.ToBeSignedUTXOInput{Address: addr1, Input: &iotago.UTXOInput{
			TransactionID:          *tx1ID,
			TransactionOutputIndex: 0,
		}}).
		AddOutput(&iotago.ExtendedOutput{Address: addr3, Amount: RequestFundsAmount}).
		Build(u.deSeriParams(), key1Signer)
	require.NoError(t, err)
	err = u.AddTransaction(spend3)
	require.Error(t, err)
}

func TestGetOutput(t *testing.T) {
	u := New()
	addr := tpkg.RandEd25519Address()
	tx, err := u.RequestFunds(addr)
	require.NoError(t, err)

	txID, err := tx.ID()
	require.NoError(t, err)

	outid0 := iotago.OutputIDFromTransactionIDAndIndex(*txID, 0)
	out0 := u.GetOutput(outid0)
	require.EqualValues(t, RequestFundsAmount, out0.Deposit())

	outid1 := iotago.OutputIDFromTransactionIDAndIndex(*txID, 1)
	out1 := u.GetOutput(outid1)
	require.EqualValues(t, u.Supply()-RequestFundsAmount, out1.Deposit())

	outidFail := iotago.OutputIDFromTransactionIDAndIndex(*txID, 5)
	require.Nil(t, u.GetOutput(outidFail))
}
