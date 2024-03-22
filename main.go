package main

import (
	"github.com/peymanzen/blockChain/blockchain"
	"github.com/peymanzen/blockChain/transactions"

	"fmt"
)

func main() {
    genesisTransactions := []transactions.Transaction{
        {Message: "Genesis Block Transaction"},
    }
    bc := blockchain.NewBlockchain(genesisTransactions)

    newTransactions := []transactions.Transaction{
        {Message: "Transaction 1"},
    }
    newBlock := blockchain.NewBlock(newTransactions, bc.Blocks[len(bc.Blocks)-1].Hash, 1)
    bc.Blocks = append(bc.Blocks, newBlock)

    fmt.Println("Blockchain initialized and a new block added!")
    for _, block := range bc.Blocks {
        fmt.Printf("Block Hash: %s\n", block.Hash)
        for _, tx := range block.Transactions {
            fmt.Printf("\tTransaction: %s\n", tx.Message)
        }
    }
}