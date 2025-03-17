package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EventSubscribe() {
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
	 */
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("key=", string(vLog.Data[:32]))
			fmt.Println("value=", string(vLog.Data[32:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics=", topics)
		}

	}
}
