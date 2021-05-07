package main

import (
	"crypto/rand"
	"github.com/binance-chain/go-sdk/common/types"

	//"flag"
	"fmt"
	"math/big"

	//"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
)

const (
	defaultBIP39Passphrase = ""
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789abcdef"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenAccount() Result {
	//privKey, _  := GenerateRandomString(64)
	//keyManager, _ := keys.NewPrivateKeyManager(privKey)

	types.Network = types.ProdNetwork
	mnemonic := "frost clock finger slender dwarf speak phone shoulder caution blossom tired nephew swift summer vast analyst potato kind maze venture very unfold pair turn"
	//if err != nil {
	//	return err
	//}
	bip44Params := keys.NewBinanceBIP44Params(0, 1)
	fmt.Printf("🔍 Finding Address end with  on %s\n", bip44Params.String())
	keyManager, _ := keys.NewMnemonicPathKeyManager(mnemonic, bip44Params.String())

	privateKey, _ := keyManager.ExportAsPrivateKey()
	addr := keyManager.GetAddr().String()
	return Result{
		privateKey,
		addr,
	}
}

type Result struct {
	privKey string
	addr    string
}

var (
	count = 0
)

func Run(c chan Result, in string) {
	var privKey string
	var addr string
	var suffix string

	for suffix != in {
		result := GenAccount()
		privKey = result.privKey
		addr = result.addr
		suffix = string([]rune(addr)[len(addr)-len(in) : len(addr)])
		count++
	}

	c <- Result{
		privKey: privKey,
		addr:    addr,
	}
}

func main() {
	// Flag
	suffix := "985"
	network := "mainnet"
	//flag.Parse()
	// Network
	//if *network == "testnet" {
	//	types.Network = types.TestNetwork
	//}

	// Loop

	fmt.Printf("🔍 Finding Address end with %s on %s\n", suffix, network)
	//for i := 0; i < runtime.NumCPU(); i++ {
	//	//fmt.Printf("🌀 Running task no. %d \n", i+1)
	//	go Run(found, suffix)
	//}

	result := GenAccount()
	//go Run(found, suffix)

	// Found
	//result := <-found

	fmt.Println("💎 ʕ•́ᴥ•̀ʔっ Got it! ")
	fmt.Println("🔒 public address :", result.addr)
	fmt.Println("🔑 private key    :", result.privKey)
}
