package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubscribeBlock() {
	fmt.Println("SubscribeBlock")
	rawurl := "wss://ethereum-sepolia-rpc.publicnode.com"
	fmt.Println("rawurl: 				", rawurl)
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是client: ", client)
	header := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), header)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case head := <-header:
			fmt.Println("head: ", head)
			block, err := client.BlockByHash(context.Background(), head.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("block: ", block.Hash().Hex())
			fmt.Println("Number: ", block.Number())
			fmt.Println("Nonce: ", block.Nonce())
			fmt.Println("Time: ", block.Time())
			fmt.Println("Transactions: ", block.Transactions())
			fmt.Println("block: ", block)
		}
	}

}
