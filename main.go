package main

import (
	"blockchain-go/blockchain"
	"fmt"
)

func main() {
	// Initialize blockchain with a genesis block
	genesisTransactions := []blockchain.Transaction{
		{Message: "Genesis Block Transaction"},
	}
	bc := blockchain.NewBlockchain(genesisTransactions)

	// Add a new block
	newTransactions := []blockchain.Transaction{
		{Message: "Transaction 1"},
	}
	newBlock := blockchain.NewBlock(newTransactions, bc.Blocks[len(bc.Blocks)-1].Hash, 1)
	bc.Blocks = append(bc.Blocks, newBlock)

	fmt.Println("Blockchain initialized and a new block added!")
}
