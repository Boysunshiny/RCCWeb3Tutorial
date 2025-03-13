package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func QueryTransaction() {

	fmt.Println("queryBalance")
	rawurl := "https://sepolia.infura.io/v3/7b7ebdc3beb1497fbb63af615e6e1cfd"
	fmt.Println("rawurl: 				", rawurl)
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID: 				", chainID)

	num := int64(7884759)
	fmt.Println("num:				", num)

	blockNumber := big.NewInt(num)

	fmt.Println("start-------------------------------BlockByNumber 通过区块号获取区块")
	block, err := client.BlockByNumber(context.Background(), blockNumber) //通过区块号获取区块
	if err != nil {
		log.Fatal(err)
	}

	for index, tx := range block.Transactions() {
		fmt.Printf("----------------%d start---------------------\n", index)
		fmt.Println("tx hash:		", tx.Hash().Hex())
		fmt.Println("tx Value:		", tx.Value().String())
		fmt.Println("tx Gas:			", tx.Gas())
		fmt.Println("tx Price:		", tx.GasPrice().Uint64())
		fmt.Println("tx Nonce:		", tx.Nonce())
		fmt.Println("tx Data:		", tx.Data())
		fmt.Println("tx To:			", tx.To().Hex())

		if sender, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
			fmt.Println("sender：		", sender.Hex())
		} else {
			fmt.Println("sender：		")
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("receipt Status:	", receipt.Status)
		fmt.Println("receipt Logs:		", receipt.Logs)
		fmt.Printf("----------------%d end---------------------\n", index)
		break
	}

	fmt.Println("end-------------------------------BlockByNumber 通过区块号获取区块")

	fmt.Println("start-------------------------------TransactionCount 通过blockHash获取block中交易的数量")
	blockHash := common.HexToHash("0xf5059972e2dc078ed0be3a67579ff70fc88fb038b5285a63ee0de67be100acfb")
	count, err := client.TransactionCount(context.Background(), blockHash) //通过blockHash获取block中交易的数量
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		fmt.Printf("----------------%d start---------------------\n", idx)
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("tx hash:		", tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Printf("----------------%d end---------------------\n", idx)
		break
	}
	fmt.Println("end-------------------------------TransactionCount 通过blockHash获取block中交易的数量")

	fmt.Println("start-------------------------------TransactionByHash 通过txHash获取交易信息")
	txHash := common.HexToHash("0x417e2e3209d8773e33f7679ef17495401fb1aa17bae367641087afb1a4bc6539")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash) //通过txHash获取交易信息
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println("tx hash:		", tx.Hash().Hex())
	fmt.Println("tx Value:		", tx.Value().String())
	fmt.Println("tx Gas:			", tx.Gas())
	fmt.Println("tx Price:		", tx.GasPrice().Uint64())
	fmt.Println("tx Nonce:		", tx.Nonce())
	fmt.Println("tx Data:		", tx.Data())
	fmt.Println("tx To:			", tx.To().Hex())
	fmt.Println("end-------------------------------TransactionByHash 通过txHash获取交易信息")
}
