package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	prefix := flag.String("prefix", "", "public address prefix")

	flag.Parse()

	hex_prefix, _ := hex.DecodeString(*prefix)

	fmt.Printf("Creating new wallet ... with prefix %x \n", hex_prefix)

	equals_prefix := false

	for !equals_prefix {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		if address[2:len(*prefix)+2] == *prefix {
			equals_prefix = true

			privateKeyBytes := crypto.FromECDSA(privateKey)
			fmt.Println("private key:", hexutil.Encode(privateKeyBytes))
			fmt.Println(address)

		}
	}
}
