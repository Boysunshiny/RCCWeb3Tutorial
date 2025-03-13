package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func SendToken() {
	fmt.Println("SendToken")
	rawurl := "https://sepolia.infura.io/v3/7b7ebdc3beb1497fbb63af615e6e1cfd"
	fmt.Println("rawurl: 				", rawurl)
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		fmt.Println("client err: 			", err)

	}
	fmt.Println("这是 clinet: 				", client)

	privateKey, err := crypto.HexToECDSA("f5ec4df35799da55cb253d9508258e36334764b95c090e643aab70ce2a2fc317")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是私钥：", privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("err")
	}
	fmt.Println("这是公钥：", publicKeyECDSA)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("这是我的账号地址：", fromAddress)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	value := big.NewInt(0)
	fmt.Println("这是ETH：", value)

	/*
		nonce
	*/
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是nonce：", nonce)

	/*
		gasPrice
	*/
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是gasPrice : ", gasPrice)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(4)) //double the gasPrice
	fmt.Println("这是gasPrice doulue: ", gasPrice)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	/*
		contractAddress
	*/
	contractAddress := common.HexToAddress("0xadB0264dE38aC757D2f98fdB5f3cCAb9a43e178f")
	fmt.Println("这是合约地址：", contractAddress)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	/*
		method ("transfer(address,uint256)")
	*/
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	fmt.Println("这是完整方法签名：", hash)
	methodID := hash.Sum(nil)[:4]
	fmt.Println("这是methodID：", methodID)
	// methodIDHex := hexutil.Encode(methodID)
	// fmt.Println("这是methodIDHex：", methodIDHex)

	toAddress := common.HexToAddress("0x79B75A8986D152bfCa0203dcAb0D45203f35CEF8") // 替换为实际的接收地址
	amount := big.NewInt(1000000000000000000)                                      // 1 ETH（以 wei 为单位）
	packedParams, err := abi.Arguments{
		{Type: mustNewType("address")},
		{Type: mustNewType("uint256")},
	}.Pack(toAddress, amount)
	if err != nil {
		log.Fatalf("Failed to pack function params: %v", err)
	}

	data := append(methodID, packedParams...)

	// 拼接函数选择器和参数

	// /*
	// 	address ("transfer(address,uint256)")
	// */
	// toAddress := common.HexToAddress("0x79B75A8986D152bfCa0203dcAb0D45203f35CEF8")
	// fmt.Println("这是接收代币：", toAddress)
	// fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	// paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32) //
	// fmt.Println("这是paddedAddress：", paddedAddress)
	// paddedAddressHex := hexutil.Encode(paddedAddress)
	// fmt.Println("这是paddedAddressHex：", paddedAddressHex)

	// /*
	// 	tokenAmount ("transfer(address,uint256)")
	// */
	// tokenAmount := new(big.Int)
	// tokenAmount.SetString("1000000000000000000", 10) // 1 tokens
	// fmt.Println("这是tokenAmount：", tokenAmount)
	// paddedTokenAmount := common.LeftPadBytes(tokenAmount.Bytes(), 32)
	// fmt.Println("这是paddedTokenAmount：", paddedTokenAmount)
	// paddedTokenAmountHex := hexutil.Encode(paddedTokenAmount)
	// fmt.Println("这是paddedTokenAmountHex：", paddedTokenAmountHex)

	// var data []byte
	// data = append(data, methodIDHex...)
	// data = append(data, paddedAddressHex...)
	// data = append(data, paddedTokenAmountHex...)
	// fmt.Println("这是data：", data)

	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	/*
		gasLimit
	*/
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	// gasLimit = 40000 //27024

	if err != nil {
		fmt.Println("gasLimit err: 			", err)
	}
	fmt.Println("这是gasLimit：", gasLimit)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	gasLimit = gasLimit * 4 //double the gasLimit
	fmt.Println("这是gasLimit doulue：", gasLimit)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	fmt.Println("--------------------发送交易-------------------")
	tx := types.NewTransaction(nonce, contractAddress, value, gasLimit, gasPrice, data)
	fmt.Println("这是tx：", tx)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是chainId：", chainId)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是signedTx：", signedTx)

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("交易发送调用完成 hx", signedTx.Hash().Hex())
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
