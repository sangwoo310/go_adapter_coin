package main
//
//import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/btcsuite/btcd/btcec"
//	"github.com/btcsuite/btcd/chaincfg"
//	"github.com/btcsuite/btcd/chaincfg/chainhash"
//	"github.com/btcsuite/btcd/txscript"
//	"github.com/btcsuite/btcd/wire"
//	"github.com/btcsuite/btcutil"
//	"math"
//	"testing"
//)
//
//func TSample(t *testing.T)  {
//
//}
//
//
////func GetPayToAddrScript(address string) []byte {
////	rcvAddress, _ := btcutil.DecodeAddress(address, &chaincfg.TestNet3Params)
////	rcvScript, _ := txscript.PayToAddrScript(rcvAddress)
////	return rcvScript
////}
//
//type TXRef struct{
//	TXHash string
//	TXOutputN int
//}
//
//// getUTXO of account which include txID, txIDIndex
//func getUTXO(address string, spendAmount int) (*[]TXRef,int, error) {
//	// get UTXO of address
//	UXTOs := []TXRef{} // get unspent Tx of an address via third party api
//	// totalInputAmount is total input amount from all txIn
//	changeAmount := totalInputAmount - spendAmount
//	return &UXTOs, changeAmount, nil
//}
//
//
//// get pub script of account
//func GetPayToAddrScript(address string) []byte {
//	rcvAddress, _ := btcutil.DecodeAddress(address, &chaincfg.TestNet3Params)
//	rcvScript, _ := txscript.PayToAddrScript(rcvAddress)
//	return rcvScript
//}
//// get private key
//func getKeyAddressFromPrivateKey(privKey string, conf *chaincfg.Params) (*btcec.PrivateKey, string, error) {
//	newWif, err := btcutil.DecodeWIF(privKey)
//	if err != nil {
//		return nil, "", err
//	}
//
//	publicKey := (*btcec.PublicKey)(&newWif.PrivKey.PublicKey).SerializeUncompressed()
//	address, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(publicKey), conf)
//	if err != nil {
//		return nil, "", err
//	}
//	return newWif.PrivKey, address.EncodeAddress(), nil
//}
//
//// debug signing function
//func Sign(rawTx, privateKey string) (string, error) {
//	var signTx wire.MsgTx
//	err := json.Unmarshal([]byte(rawTx), &signTx)
//	if err != nil {
//		return "", err
//	}
//
//	myPrivateKey, address, err := getKeyAddressFromPrivateKey(privateKey, &chaincfg.TestNet3Params)
//
//	fmt.Println("address ", address)
//	if err != nil {
//		return "", err
//	}
//	txScript := GetPayToAddrScript(address)
//	for i := 0; i < len(signTx.TxIn); i++ {
//		sig, err := txscript.SignatureScript(
//			&signTx,             // The tx to be signed.
//			i,                   // The index of the txin the signature is for.
//			txScript,            // The other half of the script from the PubKeyHash.
//			txscript.SigHashAll, // The signature flags that indicate what the sig covers.
//			myPrivateKey,        // The key to generate the signature with.
//			false)               // The compress sig flag. This saves space on the blockchain.
//
//		if err != nil {
//			return "", err
//		}
//
//		//refer from this link for sign multiple inputs https://bitcoin.stackexchange.com/questions/41209/how-to-sign-a-transaction-with-multiple-inputs
//		signTx.TxIn[i].SignatureScript = sig
//
//		//Validate signature
//		flags := txscript.StandardVerifyFlags
//		vm, err := txscript.NewEngine(txScript, &signTx, i, flags, nil, nil, signTx.TxOut[0].Value)
//		if err != nil {
//			return "", err
//		}
//		if err := vm.Execute(); err != nil {
//			return "", err
//		}
//	}
//	// return signed tx
//	return txToHex(&signTx), nil
//}
//
//
////CreateBtcRawTx create raw tx for sending btc
//func CreateRawTx(from string, to string, amount float64) (string, error) {
//	feeBitcoin := 5000
//	amountInt := amount * math.Pow(10, 8)
//	spendAmount := int(amountInt) + feeBitcoin
//	txIDs, changeAmount, _ := getUTXO(from, spendAmount)
//
//	if len(*txIDs) == 0 {
//		return "", errors.New("can not find UTXO of account")
//	}
//
//	// create new empty transaction
//	redemTx := wire.NewMsgTx(wire.TxVersion)
//	// create multiple txIns
//	for _, v := range *txIDs {
//		hash, err := chainhash.NewHashFromStr(v.TXHash)
//		if err != nil {
//			fmt.Printf("could not get hash from transaction ID: %v", err)
//			return "", err
//		}
//
//		outPoint := wire.NewOutPoint(hash, uint32(v.TXOutputN))
//		txIn := wire.NewTxIn(outPoint, nil, nil)
//		redemTx.AddTxIn(txIn)
//		fmt.Printf("TxID : %s, index %d \n", v.TXHash, v.TXOutputN)
//	}
//
//	// create TxOut
//	rcvScript := GetPayToAddrScript(to)
//	txOut := wire.NewTxOut(int64(amountInt), rcvScript)
//	redemTx.AddTxOut(txOut)
//	// create TxOut for change address, in this case, change address is sender itself
//	if changeAmount > 0 {
//		// return change BTC to its own address
//		rcvChangeAddressScript := GetPayToAddrScript(from)
//		txOut := wire.NewTxOut(int64(changeAmount), rcvChangeAddressScript)
//		redemTx.AddTxOut(txOut)
//	}
//
//	encodedTx, err := json.Marshal(redemTx)
//	if err != nil {
//		return "", err
//	}
//	// return raw tx in hex format
//	return string(encodedTx), nil
//}
