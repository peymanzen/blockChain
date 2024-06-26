package blockchain

import "github.com/peymanzen/blockchain/transactions"


type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain(genesisTransactions []transactions.Transaction) *Blockchain {
	genesisBlock := NewBlock(genesisTransactions, "", 0)
	blockchain := Blockchain{[]*Block{genesisBlock}}
	return &blockchain
}
