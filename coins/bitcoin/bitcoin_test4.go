package bitcoin

import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/golang/protobuf/proto"
	"log"
)

func wow()  {
	privKeyBytes, err := hex.DecodeString("945f423798858e24aa1ab490648013db63ad1f539ebbb8cb1399edd1d0b59716")
	if err != nil {
		log.Fatal(err)
	}
	utxoHash, err := hex.DecodeString("fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f")
	if err != nil {
		log.Fatal(err)
	}
	btcutilAddr, err := btcutil.DecodeAddress("bc1qw29x4hrt6tahz4jvuhzrq6y5el3spqt499zuay", &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	p2pkh, err := txscript.PayToAddrScript(btcutilAddr)
	if err != nil {
		log.Fatal(err)
	}
	utxo := pb.UnspentTransaction{
		OutPoint: &pb.OutPoint{
			Hash:     utxoHash,
			Index:    0,
			Sequence: 4294967295,
		},
		Amount: 625000000,
		Script: p2pkh,
	}
	plan := pb.TransactionPlan{
		Amount: 625000000,
		Utxos:  []*pb.UnspentTransaction{&utxo},
	}
	signingInput := pb.SigningInput{
		HashType:      1, // 1 for TWBitcoinSigHashTypeAll
		Amount:        625000000,
		ByteFee:       1,
		ToAddress:     "1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
		ChangeAddress: "1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU",
		PrivateKey:    [][]byte{privKeyBytes},
		Utxo:          []*pb.UnspentTransaction{&utxo},
		CoinType:      0,
		Plan:          &plan,
	}
	inputBytes, err := proto.Marshal(&signingInput)
	if err != nil {
		log.Fatal(err)
	}
	inputData := TWDataCreateWithGoBytes(inputBytes)
	defer C.TWDataDelete(inputData)
	outputData := C.TWAnySignerSign(inputData, C.TWCoinTypeBitcoin)
	var output pb.SigningOutput
	err = proto.Unmarshal(TWDataGoBytes(outputData), &output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("<== bitcoin signed output: %x\n", output.Encoded)

}
