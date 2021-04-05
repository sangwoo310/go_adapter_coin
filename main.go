package main

import (
	"fmt"
	"go_blockchain/factory"
	"log"
)

func main() {
	fmt.Println("This is the blockchain")

	c, err := factory.GetCoinInstance("btc")
	if err != nil {
		log.Fatalln(err)
	}

	c.PublicKeyToAddress()
	c.Sign()
}
