package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
	"github.com/peymanzen/blockchain/transactions"
)
 
type Block struct {
	sequenceNumber int64
	PrevBlockHash  string
	Timestamp      int64
	Transactions   []transactions.Transaction
	Hash           string
	ProposerID     int
}

func NewBlock(transactions []transactions.Transaction, prevBlockHash string, proposerID int) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		ProposerID:    proposerID,
	}
	block.sequenceNumber = block.Timestamp
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	input := strconv.FormatInt(b.sequenceNumber, 10) + b.PrevBlockHash + strconv.FormatInt(b.Timestamp, 10) + strconv.Itoa(b.ProposerID)
	for _, tx := range b.Transactions {
		input += tx.Message
	}
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
