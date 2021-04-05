package bitcoin

import "go_blockchain/coins"

type Bitcoin struct {
	CoinName 		string
	Algorithm 		string
	AddressType 	string
}

func New(coinName string) (coins.ICoin, error) {
	s := Bitcoin{
		CoinName: coinName,
	}

	return s, nil
}

func (c Bitcoin) PublicKeyToAddress() error {
	return nil
}

func (c Bitcoin) Sign() error {
	return nil
}

func getKeyPair() {

}