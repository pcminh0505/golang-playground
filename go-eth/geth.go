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

var infuraURL = "https://mainnet.infura.io/v3/c9f3edf2c9a340cbaa7a96058862910b"

func main () {
	client, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}

	fmt.Println("Block number:",block.Number())

	addr := "0x73bceb1cd57c711feac4224d062b0f6ff338501e"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get balance: %v", err)
	}

	fmt.Println("Balance:", balance)
	// 1 ETH = 10^18 wei
	fBalance := new (big.Float)
	fBalance.SetString(balance.String())
	fmt.Println(fBalance)

	balanceETH := new (big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Balance in ETH:", balanceETH)
}