package consensus

import "github.com/peymanzen/blockchain/blockchain"

type PBFT struct {
	NodeID    int
	Blockchain *blockchain.Blockchain
	CurrentSequence int64
	PrePrepares map[int64]Message
	Prepares    map[int64][]Message
	Commits     map[int64][]Message

}

func NewPBFT(id int, bc *blockchain.Blockchain) *PBFT {
    return &PBFT{
        NodeID :         id,
        Blockchain: bc,
		PrePrepares: make(map[int64]Message),
		Prepares:    make(map[int64][]Message),
		Commits:     make(map[int64][]Message),
    }
}


// Placeholder methods for handling consensus stages.
func (pbft *PBFT) HandlePrePrepare(msg Message) {
	pbft.PrePrepares[msg.SequenceID] = msg
    // Implement handling of PrePrepare messages here.
}

func (pbft *PBFT) HandlePrepare(msg Message) {
	pbft.Prepares[msg.SequenceID] = append(pbft.Prepares[msg.SequenceID], msg)
    // Implement handling of Prepare messages here.
}

func (pbft *PBFT) HandleCommit(msg Message) {
	pbft.Commits[msg.SequenceID] = append(pbft.Commits[msg.SequenceID], msg)
    // Implement handling of Commit messages here.
}