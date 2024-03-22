package consensus

import "github.com/peymanzen/blockChain/blockchain"

type PBFT struct {
	ID    int
	Blockchain *blockchain.Blockchain

}

func NewPBFT(id int, bc *blockchain.Blockchain) *PBFT {
    return &PBFT{
        ID:         id,
        Blockchain: bc,
    }
}


// Placeholder methods for handling consensus stages.
func (pbft *PBFT) HandlePrePrepare(msg Message) {
    // Implement handling of PrePrepare messages here.
}

func (pbft *PBFT) HandlePrepare(msg Message) {
    // Implement handling of Prepare messages here.
}

func (pbft *PBFT) HandleCommit(msg Message) {
    // Implement handling of Commit messages here.
}