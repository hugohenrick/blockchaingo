package main

import (
	"fmt"

	blockchain "github.com/hugohenrick/blockchaingo/blockchain"
)

func main() {
	myBlockChainAddress := "my_blockchain_address"
	blockchain := blockchain.NewBlockchain(myBlockChainAddress)
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("my %.1f\n", blockchain.CalulateTotalAmount(myBlockChainAddress))
	fmt.Printf("C %.1f\n", blockchain.CalulateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockchain.CalulateTotalAmount("D"))
}
