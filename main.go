package main

import (
	"fmt"
	"github.com/binance-chain/go-sdk/keys"
)

type Account struct {
	PrivateKey string
	addr       string
}

func GenAccount() Account {
	km, _ := keys.NewKeyManager()

	privateKey, _ := km.ExportAsPrivateKey()
	addr := km.GetAddr().String()

	return Account{
		privateKey,
		addr,
	}
}

func main() {
	result := GenAccount()
	fmt.Println("🔒 public address :", result.addr)
	fmt.Println("🔑 private key    :", result.PrivateKey)
}
