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

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/4W0HnZRWnhvYz9m5RDjfwmxIFt9tFKNB")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x550cb21D24831A7CC203c7a72CD5C27D6791D99B")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balanceAt) // 25729324269165216042
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pendingBalance) // 25729324269165216042

}
