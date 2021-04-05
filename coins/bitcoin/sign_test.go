package bitcoin

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"testing"
)

func Test_CreateRedeemTx(t *testing.T) {
	txInput := wire.TxIn{
		PreviousOutPoint: wire.OutPoint{},
		SignatureScript:  nil,
		Witness:          nil,
		Sequence:         0,
	}

	txOutput := wire.TxOut{
		Value:    0,
		PkScript: nil,
	}

	out_1 := wire.NewOutPoint()

	txInput_1 := wire.NewTxIn(&chaincfg.)

	redeemTx := wire.NewMsgTx(wire.TxVersion)

	redeemTx.AddTxIn()
	redeemTx.AddTxOut()
}