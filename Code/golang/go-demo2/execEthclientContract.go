package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ExecEthclientContract() {

	// 连接以太坊节点
	url := "https://eth-sepolia.g.alchemy.com/v2/k1Q_rwvQKEyIFAenW5p_mTSmkvhYPACP"
	client, err := ethclient.Dial(url)
	if err != nil {

	}
	fmt.Println(client)
	privateKey, err := crypto.HexToECDSA("f5ec4df35799da55cb253d9508258e36334764b95c090e643aab70ce2a2fc317")
	if err != nil {
		log.Fatal(err)

	}
	// fmt.Println("这是私钥			", privateKey)
	pubKey := privateKey.Public()
	// fmt.Println("这是公钥			", pubKey)
	publicKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// fmt.Println("这是公钥			", publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("这是地址			", address)

	data, err := os.ReadFile("./Store_sol_Store.abi")
	if err != nil {
		log.Fatal(err)
	}
	json := string(data)
	fmt.Println("这是字符串ABI			", json)
	contractAbi, err := abi.JSON(strings.NewReader(json))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("这是ContractABI			", contractAbi)
	if err != nil {
		log.Fatal(err)
	}
	/////////////////////////////////////////////////////////////////////
	/*
		// 调用合约方法 setItem
	*/
	methodName := "setItem"
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key_use_abi"))
	copy(value[:], []byte("demo_save_value_use_abi_11111"))
	input, err := contractAbi.Pack(methodName, key, value)

	if err != nil {
		log.Fatal(err)

	}
	/*
		gasPrice
	*/
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	/*
		nonce
	*/
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是nonce			", nonce)
	/*
		chainID
	*/
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是chainID			", chainID)

	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 300000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是signedTx			", signedTx)

	client.SendTransaction(context.Background(), signedTx)

	waitForTransactionReceipt(context.Background(), client, signedTx.Hash())

	/////////////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////////
	/*

		获取合约方法返回值
	*/

	// fmt.Println("这是input			", input)

	// 查询刚刚设置的值
	callInput, err := contractAbi.Pack("items", key)
	if err != nil {
		log.Fatal(err)
	}

	to := common.HexToAddress(contractAddr)
	callOpts := ethereum.CallMsg{
		To:   &to,       // &common.Address{0x0000000000000000000000000000000000000123},
		Data: callInput, // common.FromHex("0x12345678"),
	}

	result, err := client.CallContract(context.Background(), callOpts, nil)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked [32]byte
	contractAbi.UnpackIntoInterface(&unpacked, "items", result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("结果", unpacked == value)

	/////////////////////////////////////////////////////////////////////

}
