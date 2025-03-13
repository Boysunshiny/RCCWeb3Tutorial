package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GenerateKey() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("original private key:", privateKey)
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key in bytes:", privateKeyBytes)
	privateKeyHexStr := hexutil.Encode(privateKeyBytes[2:]) // 去掉'0x'
	fmt.Println("private key in hexstr:", privateKeyHexStr)

	publicKey := privateKey.Public()
	fmt.Println("public key:", publicKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println("address:", address)

	hash := sha3.NewLegacyKeccak256()
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("public key in bytes:", publicKeyBytes)
	// fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes[1:])) // 去掉第一个byte
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println("public key hash:", hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位

}
