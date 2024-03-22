package consensus


import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)
type MessageType int


const (
    PrePrepare MessageType = iota
	Prepare
	Commit
)

type Message struct {
    Type       MessageType
	BlockHash  string
	SequenceID int64
	Timestamp  int64
	SenderID   int
	Signature  string
}

func SignMessage(msg *Message, privateKey string) {
	input := msg.BlockHash + string(msg.SequenceID) + string(msg.Timestamp)+privateKey
	hash := sha256.Sum256([]byte(input))
	msg.Signature = hex.EncodeToString(hash[:])
}

func NewMessage(msgType MessageType, blockHash string, sequenceID int64, senderID int) Message {
	msg := Message{
		Type:       msgType,
		BlockHash:  blockHash,
		SequenceID: sequenceID,
		Timestamp:  time.Now().Unix(),
		SenderID:   senderID,
	}
	// Sign the message with a 'private key', simplified here
	SignMessage(&msg, "nodePrivateKey")
	return msg
}