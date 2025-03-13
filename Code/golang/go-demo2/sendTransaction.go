package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendTransaction() {
	rawurl := "https://sepolia.infura.io/v3/7b7ebdc3beb1497fbb63af615e6e1cfd"
	fmt.Println("rawurl: 				", rawurl)
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("f5ec4df35799da55cb253d9508258e36334764b95c090e643aab70ce2a2fc317")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是私钥：", privateKey)
	publicKey := privateKey.Public()
	fmt.Println("这是公钥：", privateKey)
	publickeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publickeyECDSA)
	fmt.Println("这是地址：", fromAddress)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是帐号的nonce: ", nonce)

	value := big.NewInt(1000000000000000) // in wei (1 eth)\
	fmt.Println("这是value: ", value)
	gasLimit := uint64(21000)
	fmt.Println("这是gasLimit: ", gasLimit)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("这是gasPrice: ", gasPrice)

	toAddress := common.HexToAddress("0x79B75A8986D152bfCa0203dcAb0D45203f35CEF8")

	fmt.Println("这是toAddress: ", toAddress)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	fmt.Println("这是tx: ", tx)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是chainId: ", chainId)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是signed: ", signedTx.Hash().Hex())

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("发送交易完成tx：", signedTx.Hash().Hex())
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
