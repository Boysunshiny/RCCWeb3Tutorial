package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-demo2/store"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AbigenDeploy() {

	fmt.Println("abigenDeploy")

	// rawurl := "https://sepolia.infura.io/v3/7b7ebdc3beb1497fbb63af615e6e1cfd"
	// rawurl := "http://127.0.0.1:8545"
	rawurl := "https://eth-sepolia.g.alchemy.com/v2/k1Q_rwvQKEyIFAenW5p_mTSmkvhYPACP"

	fmt.Println("rawurl: 				", rawurl)
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("client: 				", client)

	privateKey, err := crypto.HexToECDSA("f5ec4df35799da55cb253d9508258e36334764b95c090e643aab70ce2a2fc317")

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("这是私钥：", privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("这是账号地址：", fromAddress)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("这是nonce: 				", nonce)

	value := big.NewInt(0) // in wei (1 eth)
	fmt.Println("这是金额value: 				", value)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// gasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
	fmt.Println("这是gasPrice 				", gasPrice)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是chainId: 				", chainId)

	/*
		准备部署
	*/
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是auth: 				", auth)

	auth.Nonce = big.NewInt(int64(nonce))
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	auth.Value = value
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	auth.GasLimit = uint64(300000)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	auth.GasPrice = gasPrice
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	/*
	 部署
	*/

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("这是地址：", address.Hex())
	fmt.Println("这是交易：", tx.Hash().Hex())
	fmt.Println("这是实例：", instance)
}
