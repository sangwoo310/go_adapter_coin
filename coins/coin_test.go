package coins

import (
	"github.com/stretchr/testify/assert"
	"go_blockchain/coins/bitcoin"
	"testing"
)

func Test_Interface(t *testing.T) {
	c, err := bitcoin.New("btc")
	if err != nil {
		assert.NoError(t, err)
	}

	c.PublicKeyToAddress()
	c.Sign()
}