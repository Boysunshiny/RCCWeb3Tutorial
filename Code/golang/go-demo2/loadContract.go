package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-demo2/store"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddr = "0xb6f840199A4e1d7d70561126B61B29665D8d2898"
)

func LoadContract() {

	fmt.Println("LoadContract")

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
	fmt.Println("这是auth.Nonce: 				", auth.Nonce)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	auth.Value = value
	fmt.Println("这是auth.Value: 				", auth.Value)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	auth.GasLimit = uint64(300000)
	fmt.Println("这是auth.GasLimit: 				", auth.GasLimit)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	auth.GasPrice = gasPrice
	fmt.Println("这是auth.GasPrice: 				", auth.GasPrice)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	key := [32]byte{}
	val := [32]byte{}

	copy(key[:], []byte("hello")) // 将字符串转换为 [32]byte
	copy(val[:], []byte("world")) // 将字符串转换为 [32]byte

	/*
		加载合约
	*/
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	/*
		调用合约方法
	*/
	tx, err := storeContract.SetItem(auth, key, val)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx: ", tx.Hash().Hex())

}
