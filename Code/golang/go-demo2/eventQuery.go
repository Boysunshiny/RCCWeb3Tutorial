package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EventQuery() {
	/*
	   client
	*/
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/k1Q_rwvQKEyIFAenW5p_mTSmkvhYPACP")
	if err != nil {
		log.Fatal(err)

	}

	/*
		contractAddress
	*/
	contractAddress := common.HexToAddress(contractAddr)

	/*
	   query
	*/
	query := ethereum.FilterQuery{
		// FromBlock: big.NewInt(6920583),
		// ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
		// Topics: [][]common.Hash{
		//  {},
		//  {},
		// },
	}
	/*
		filter
	*/
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for _, vLog := range logs {

		fmt.Println("key=", string(vLog.Data[:32]))
		fmt.Println("value=", string(vLog.Data[32:]))

		var topics []string
		for i := range vLog.Topics {
			topics = append(topics, vLog.Topics[i].Hex())
		}
		fmt.Println("topics[0]=", topics[0])
		if len(topics) > 1 {
			fmt.Println("indexed topics:", topics[1:])
		}
	}

}
