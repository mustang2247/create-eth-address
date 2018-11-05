package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"regexp"
)

func main() {
	//var address = "0x888124fA295A0220Ef2e1749D0065936F6471ac1"
	//Check(address)
	//
	//client,err := connectToRpc()
	//fmt.Println("end connecting")
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//getBalance(client, address)
	//
	//getFullBlock(client, 3030902)
	//


	client, err := ethclient.Dial("wss://mainnet.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			//log.Panic(err)
			fmt.Sprint(err)
			continue
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				//log.Panic(err)
				fmt.Sprint(err)
				continue
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time().Uint64())     // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}

}

func getFullBlock(client *ethclient.Client, blockNum int64) {
	blockNumber := big.NewInt(blockNum)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
			fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}

	//fmt.Println(block.Number().Uint64())     // 5671744
	//fmt.Println(block.Time().Uint64())       // 1527211625
	//fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	//fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	//fmt.Println(len(block.Transactions()))   // 144
}

//特别注意，这里的address就是你要查询的以太坊余额的地址。一般是0xddddd 这样的形式
//特别注意的是以太坊的Decimal是18，那么我们获得的余额要乘以10^-18,才能得到正常的以太坊数量。
//以太坊的其他token也一样，会有不同的Decimal，但是会有相应的方法获得，这个不需要担心，在下一个连载会讲到。
func getBalance(client *ethclient.Client, address string) {
	account := common.HexToAddress(address)
	//fmt.Println(account.Hex())

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		fmt.Println(err.Error())
		//beego.Error(err.Error())
	} else {

		//fmt.Println(balance)
		//这个就是地址中以太坊的余额
		balanceV := float64(balance.Int64()) * math.Pow(10, -18)
		fmt.Println(balanceV)
	}

	//blockInfo,err := client.BlockByNumber(context.Background(), big.NewInt(6376566))
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(blockInfo.Transactions().Len())
}

/**
	连接以太坊
 */
func connectToRpc() (*ethclient.Client, error) {
	//client, err := ethclient.Dial("https://mainnet.infura.io/OGATmNvVTVWjw3Lu9e9i")
	//client, err := ethclient.Dial("http://35.236.141.214:38845")
	client, err := ethclient.Dial("https://rinkeby.infura.io/OGATmNvVTVWjw3Lu9e9i")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	//conn := ethclient.NewClient(client)
	return client, err
}

func Check(address string) {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString(address)) // is valid: true
	//fmt.Printf("is valid: %v\n", re.MatchString(address)) // is valid: false

	//client, err := ethclient.Dial("https://mainnet.infura.io")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 0x Protocol Token (ZRX) smart contract address
	//address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	//bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//isContract := len(bytecode) > 0
	//
	//fmt.Printf("is contract: %v\n", isContract) // is contract: true
	//
	//// a random user account address
	//address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	//bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//isContract = len(bytecode) > 0
	//
	//fmt.Printf("is contract: %v\n", isContract) // is contract: false
}