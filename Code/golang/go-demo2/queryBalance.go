package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// queryBalance
// 查询余额
// 查询余额

func QueryBalance() {
	//https://eth-sepolia.g.alchemy.com/v2/4W0HnZRWnhvYz9m5RDjfwmxIFt9tFKNB
	//https://sepolia.infura.io/v3/7b7ebdc3beb1497fbb63af615e6e1cfd

	fmt.Println("queryBalance")
	rawurl := "https://sepolia.infura.io/v3/7b7ebdc3beb1497fbb63af615e6e1cfd"
	fmt.Println("rawurl: 				", rawurl)
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	acc := "0x568e17144037b15833531BEB88743fdc9eC776aA"
	fmt.Println("acc:				", acc)
	account := common.HexToAddress(acc)
	balance, err := client.BalanceAt(context.Background(), account, nil) //balance
	if err != nil {
		log.Fatal(err)
	}

	qblock, err := client.BlockNumber(context.Background()) //balance
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block:				", qblock)
	fmt.Println("balance:			", balance)
	fmt.Println("-------------------------------------")
	num := int64(7884759)
	fmt.Println("num:				", num)
	blockNumber := big.NewInt(num)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber) //balance
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance:			", balanceAt)
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue:			", ethValue)

	fmt.Println("-------------------------------------")
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account) //balance
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pendingBalance:			", pendingBalance)

}
