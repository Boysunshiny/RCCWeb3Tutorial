package main

import (
	"context"
	"fmt"
	"go-demo2/store"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddress = "0xb6f840199A4e1d7d70561126B61B29665D8d2898"
)

func ExecAbigenContract() {

	/*
		ad
	*/
	fmt.Println("ExecAbigenContract")
	rawurl := "https://eth-sepolia.g.alchemy.com/v2/k1Q_rwvQKEyIFAenW5p_mTSmkvhYPACP"
	fmt.Println("rawurl: 				", rawurl)
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接网络", client)

	privateKey, err := crypto.HexToECDSA("f5ec4df35799da55cb253d9508258e36334764b95c090e643aab70ce2a2fc317")

	if err != nil {
		log.Fatal(err)
	}
	chainId, err := client.NetworkID(context.Background())
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("chainId: ", chainId)

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("opts: ", opts)

	address := common.HexToAddress(contractAddress)
	storeContract, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("加载实例的合约: ", storeContract)

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], "11111")
	copy(value[:], "22222")
	fmt.Println("key: ", hexutil.Encode(key[:]))
	fmt.Println("value: ", hexutil.Encode(value[:]))
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	/*
		调用 合约
		function setItem(bytes32 key, bytes32 value) external {
			items[key] = value;
			emit ItemSet(key, value);
		}
	*/
	tx, err := storeContract.SetItem(opts, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("调用合约方法SetItem 的交易哈希:", tx.Hash())

	/*
		获取 合约
		items(bytes32)
		//mapping(bytes32 => bytes32) public items;
	*/
	callOpts := bind.CallOpts{Context: context.Background()}
	result, err := storeContract.Items(&callOpts, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("获取合约属性mapping的值:", result)

	fmt.Println("设置与获取的结果对比 ", result == value)

}
