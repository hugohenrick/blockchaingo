package main

import (
	"fmt"
	"log"

	"github.com/hugohenrick/blockchaingo/blockchain"
	"github.com/hugohenrick/blockchaingo/wallet"
)

func intit() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet
	t := wallet.NewTranascation(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0)

	// Blockchain
	blockchain := blockchain.NewBlockchain(walletM.BlockChainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0, walletA.PublicKey(), t.GenerateSiginature())

	fmt.Println("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("Balance A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockChainAddress()))
	fmt.Printf("Balance B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockChainAddress()))
	fmt.Printf("Balance M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockChainAddress()))
}
