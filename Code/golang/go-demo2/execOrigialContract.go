package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ExecOrigialContract() {
	/*
		client
	*/
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/k1Q_rwvQKEyIFAenW5p_mTSmkvhYPACP")
	if err != nil {
		log.Fatal(err)
	}
	/*
		privateKey
	*/
	privateKey, err := crypto.HexToECDSA("f5ec4df35799da55cb253d9508258e36334764b95c090e643aab70ce2a2fc317")
	if err != nil {
		log.Fatal(err)
	}
	/*
		publicKey
	*/
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	/*
		address
	*/
	address := crypto.PubkeyToAddress(*publicKey)

	/*
		nonce
	*/
	nonce, err := client.PendingNonceAt(context.Background(), address)
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
		chainId
	*/

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)

	}

	/*
		data
	*/

	methodSignature := []byte("setItem(bytes32,bytes32)")
	methodSelector := crypto.Keccak256(methodSignature)[:4]

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("kkk"))
	copy(value[:], []byte("vvv"))

	/////////////////////////////////////////
	var data []byte
	data = append(data, methodSelector...)
	data = append(data, key[:]...)
	data = append(data, value[:]...)
	fmt.Println("data ", hexutil.Encode(data))
	/////////////////////////////////////////

	/////////////////////////////////////////
	file, err := os.ReadFile("./Store_sol_Store.abi")
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		log.Fatal(err)
	}
	pack, err := contractAbi.Pack("setItem", key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pack ", hexutil.Encode(pack))
	/////////////////////////////////////////
	/*
		gasLimit
	*/
	to := common.HexToAddress(contractAddr)
	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To:   &to,
	// 	Data: data,
	// })
	// if err != nil {
	// 	log.Fatal(err)

	// }
	//直接获取值

	tx := types.NewTransaction(nonce, to, big.NewInt(0), 300000, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	waitForTransactionReceipt(context.Background(), client, signedTx.Hash())

	fmt.Printf("tx success: %s\n", signedTx.Hash().Hex())

	// tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 300000, gasPrice, pack)
	// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("tx success: %s\n", signedTx.Hash().Hex())
	// client.SendTransaction(context.Background(), signedTx)

	// waitForTransactionReceipt(context.Background(), client, signedTx.Hash())

}
