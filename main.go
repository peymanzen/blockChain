// package main

// import (
// 	"github.com/peymanzen/blockChain/blockchain"
// 	"github.com/peymanzen/blockChain/transactions"

// 	"fmt"
// )

// func main() {
//     genesisTransactions := []transactions.Transaction{
//         {Message: "Genesis Block Transaction"},
//     }
//     bc := blockchain.NewBlockchain(genesisTransactions)

//     newTransactions := []transactions.Transaction{
//         {Message: "Transaction 1"},
//     }
//     newBlock := blockchain.NewBlock(newTransactions, bc.Blocks[len(bc.Blocks)-1].Hash, 1)
//     bc.Blocks = append(bc.Blocks, newBlock)

//     fmt.Println("Blockchain initialized and a new block added!")
//     for _, block := range bc.Blocks {
//         fmt.Printf("Block Hash: %s\n", block.Hash)
//         for _, tx := range block.Transactions {
//             fmt.Printf("\tTransaction: %s\n", tx.Message)
//         }
//     }
// }
package main

import (
    "github.com/peymanzen/blockchain/blockChain"
    "github.com/peymanzen/blockchain/consensus"
    "github.com/peymanzen/blockchain/networking"
	"github.com/peymanzen/blockchain/transactions"
    "fmt"
    "os"
	"net/http"
)

func main() {
    // Node configuration - this would be more dynamic in a real application
    nodeID := 1
    port := 8080

    // Initialize the blockchain
    bc := blockchain.NewBlockchain([]transactions.Transaction{{Message: "Genesis Block"}})

    // Initialize PBFT
    pbft := consensus.NewPBFT(nodeID, bc)

    // Start the server to listen for incoming messages
    messageHandler := func(w http.ResponseWriter, r *http.Request) {
        // Here you'd add logic to handle incoming PBFT messages
        fmt.Println("Received a message")
    }
    go networking.StartServer(port, messageHandler)

    fmt.Printf("Node %d running on port %d\n", nodeID, port)

    // Placeholder for sending a message - in practice, this would be part of the PBFT logic
    if nodeID == 1 {
        err := networking.SendMessage("localhost:8081", []byte("Test message"))
        if err != nil {
            fmt.Printf("Error sending message: %s\n", err)
            os.Exit(1)
        }
    }
}