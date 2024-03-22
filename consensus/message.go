package consensus

type MessageType int


const (
    PrePrepare MessageType = iota
	Prepare
	Commit
)

type Message struct {
    Type MessageType
	BlockHash string
	SenderID int
    SequenceID int64
}