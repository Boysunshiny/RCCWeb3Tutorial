package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ethereumDeployCode = "608060405234801561000f575f80fd5b5060405161087538038061087583398181016040528101906100319190610193565b805f908161003f91906103e7565b50506104b6565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a58261005f565b810181811067ffffffffffffffff821117156100c4576100c361006f565b5b80604052505050565b5f6100d6610046565b90506100e2828261009c565b919050565b5f67ffffffffffffffff8211156101015761010061006f565b5b61010a8261005f565b9050602081019050919050565b8281835e5f83830152505050565b5f610137610132846100e7565b6100cd565b9050828152602081018484840111156101535761015261005b565b5b61015e848285610117565b509392505050565b5f82601f83011261017a57610179610057565b5b815161018a848260208601610125565b91505092915050565b5f602082840312156101a8576101a761004f565b5b5f82015167ffffffffffffffff8111156101c5576101c4610053565b5b6101d184828501610166565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022857607f821691505b60208210810361023b5761023a6101e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610262565b6102a78683610262565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102eb6102e66102e1846102bf565b6102c8565b6102bf565b9050919050565b5f819050919050565b610304836102d1565b610318610310826102f2565b84845461026e565b825550505050565b5f90565b61032c610320565b6103378184846102fb565b505050565b5b8181101561035a5761034f5f82610324565b60018101905061033d565b5050565b601f82111561039f5761037081610241565b61037984610253565b81016020851015610388578190505b61039c61039485610253565b83018261033c565b50505b505050565b5f82821c905092915050565b5f6103bf5f19846008026103a4565b1980831691505092915050565b5f6103d783836103b0565b9150826002028217905092915050565b6103f0826101da565b67ffffffffffffffff8111156104095761040861006f565b5b6104138254610211565b61041e82828561035e565b5f60209050601f83116001811461044f575f841561043d578287015190505b61044785826103cc565b8655506104ae565b601f19841661045d86610241565b5f5b828110156104845784890151825560018201915060208501945060208101905061045f565b868310156104a1578489015161049d601f8916826103b0565b8355505b6001600288020188555050505b505050505050565b6103b2806104c35f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806348f343f31461004357806354fd4d5014610073578063f56256c714610091575b5f80fd5b61005d600480360381019061005891906101d7565b6100ad565b60405161006a9190610211565b60405180910390f35b61007b6100c2565b604051610088919061029a565b60405180910390f35b6100ab60048036038101906100a691906102ba565b61014d565b005b6001602052805f5260405f205f915090505481565b5f80546100ce90610325565b80601f01602080910402602001604051908101604052809291908181526020018280546100fa90610325565b80156101455780601f1061011c57610100808354040283529160200191610145565b820191905f5260205f20905b81548152906001019060200180831161012857829003601f168201915b505050505081565b8060015f8481526020019081526020015f20819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d48282604051610194929190610355565b60405180910390a15050565b5f80fd5b5f819050919050565b6101b6816101a4565b81146101c0575f80fd5b50565b5f813590506101d1816101ad565b92915050565b5f602082840312156101ec576101eb6101a0565b5b5f6101f9848285016101c3565b91505092915050565b61020b816101a4565b82525050565b5f6020820190506102245f830184610202565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61026c8261022a565b6102768185610234565b9350610286818560208601610244565b61028f81610252565b840191505092915050565b5f6020820190508181035f8301526102b28184610262565b905092915050565b5f80604083850312156102d0576102cf6101a0565b5b5f6102dd858286016101c3565b92505060206102ee858286016101c3565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061033c57607f821691505b60208210810361034f5761034e6102f8565b5b50919050565b5f6040820190506103685f830185610202565b6103756020830184610202565b939250505056fea26469706673582212205aae308f77654b000c9d222eff2d9f2bd2ac18d990b10774842e4309d4e3e15664736f6c634300081a0033"

func EthereumDeploy() {

	fmt.Println("EthereumDeploy")

	// rawurl := "https://sepolia.infura.io/v3/7b7ebdc3beb1497fbb63af615e6e1cfd"
	// rawurl := "http://127.0.0.1:8545"
	rawurl := "https://eth-sepolia.g.alchemy.com/v2/k1Q_rwvQKEyIFAenW5p_mTSmkvhYPACP"

	fmt.Println("rawurl: 				", rawurl)
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("client: 				", client)

	privateKey, err := crypto.HexToECDSA("f5ec4df35799da55cb253d9508258e36334764b95c090e643aab70ce2a2fc317")

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("这是私钥：", privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("这是账号地址：", fromAddress)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("这是nonce: 				", nonce)

	value := big.NewInt(0) // in wei (1 eth)
	fmt.Println("这是金额value: 				", value)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// gasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
	fmt.Println("这是gasPrice 				", gasPrice)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是chainId: 				", chainId)

	/*
		准备部署
	*/
	// 构造部署数据
	// 1. 获取合约字节码
	bytecode := "0x" + ethereumDeployCode // 替换为实际的合约字节码

	// 2. 编码构造函数参数
	input := "1.0" // 字符串类型的参数
	packedParams, err := abi.Arguments{
		{Type: mustNewType("string")},
	}.Pack(input)

	if err != nil {
		log.Fatalf("Failed to pack constructor params: %v", err)
	}
	// 3. 将字节码和参数拼接
	data := append(common.FromHex(bytecode), packedParams...)

	gasLimit := uint64(300000)

	tx := types.NewContractCreation(nonce, value, gasLimit, gasPrice, data)

	fmt.Println("这是tx: 				", tx)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainId), privateKey)

	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是signed: 				", signedTx.Hash().Hex())

	recipt, err := waitForTransactionReceipt(context.Background(), client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("这是recipt: 				", recipt.ContractAddress.Hex())

}

func waitForTransactionReceipt(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		recipt, err := client.TransactionReceipt(context.Background(), txHash)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if err != nil {
				fmt.Println("Waiting for transaction receipt...")
			} else {
				return recipt, nil
			}
			time.Sleep(10 * time.Second)
		}
	}
}
