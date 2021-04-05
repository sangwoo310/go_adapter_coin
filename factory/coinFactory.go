package factory

import (
	"errors"
	"go_blockchain/coins"
	"go_blockchain/coins/bitcoin"
)

const (
	btc = "bitcoin"
	eth = "ethereum"
)

func GetCoinInstance(mainNetId string) (coins.ICoin, error) {
	switch mainNetId {
	case btc:
		return bitcoin.New(btc)
	}

	return nil, errors.New("cannot set coin instance")
}