package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// 主函数
func main() {

	// go run main.go queryBalance.go queryTransaction.go queryRecipt.go generateKey.go sendTransaction.go sendToken.go subscribeBlock.go abigenDeploy.go ethereumDeploy.go loadContract.go execAbigenContract.go execEthclientContract.go execOrigialContract.go eventQuery.go
	testid := 14
	switch testid {
	case 1:
		fmt.Println("开始----》	QueryBalance")
		QueryBalance()
		fmt.Println("开始----》	QueryBalance")
	case 2:

		fmt.Println("开始----》	QueryTransaction")
		QueryTransaction()
		fmt.Println("开始----》	QueryTransaction")
	case 3:
		fmt.Println("开始----》	QueyrRecipt")
		QueyrRecipt()
		fmt.Println("开始----》	QueyrRecipt")
	case 4:
		fmt.Println("开始----》	GenerateKey")
		GenerateKey()
		fmt.Println("开始----》	GenerateKey")
	case 5:
		fmt.Println("开始----》	SendTransaction")
		SendTransaction()
		fmt.Println("开始----》	SendTransaction")
	case 6:
		fmt.Println("开始----》	SendToken")
		SendToken()
		fmt.Println("开始----》	SendToken")
	case 7:
		fmt.Println("开始----》	SubscribeBlock")
		SubscribeBlock()
		fmt.Println("开始----》	SubscribeBlock")

	case 8:
		fmt.Println("开始----》	AbigenDeploy")
		AbigenDeploy()
		fmt.Println("开始----》	AbigenDeploy")

	case 9:
		fmt.Println("开始----》	EthereumDeploy")
		EthereumDeploy()
		fmt.Println("开始----》	EthereumDeploy")

	case 10:
		fmt.Println("开始----》	LoadContract")
		LoadContract()
		fmt.Println("开始----》	LoadContract")

	case 11:
		fmt.Println("开始----》	ExecAbigenContract")
		ExecAbigenContract()
		fmt.Println("开始----》	ExecAbigenContract")

	case 12:
		fmt.Println("开始----》	ExecEthclientContract")
		ExecEthclientContract()
		fmt.Println("开始----》	ExecEthclientContract")

	case 13:
		fmt.Println("开始----》	ExecOrigialContract")
		ExecOrigialContract()
		fmt.Println("开始----》	ExecOrigialContract")
	case 14:
		fmt.Println("开始----》	EventQuery")
		EventQuery()
		fmt.Println("开始----》	EventQuery")
	}

}

// 辅助函数：创建 abi.Type
func mustNewType(typeStr string) abi.Type {
	typ, err := abi.NewType(typeStr, "", nil)
	if err != nil {
		panic(fmt.Sprintf("Failed to create type: %v", err))
	}
	return typ
}
