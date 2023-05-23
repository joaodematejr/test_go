package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.PrevHash
	for _, transaction := range block.Transactions {
		record += transaction.Sender + transaction.Receiver + fmt.Sprintf("%.2f", transaction.Amount)
	}
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, transactions []Transaction) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Transactions = transactions
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func isChainValid(chain []Block) bool {
	for i := 1; i < len(chain); i++ {
		currentBlock := chain[i]
		previousBlock := chain[i-1]

		if currentBlock.Hash != calculateHash(currentBlock) {
			return false
		}

		if currentBlock.PrevHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func main() {
	genesisBlock := Block{}
	genesisBlock.Index = 0
	genesisBlock.Timestamp = time.Now().String()
	genesisBlock.Transactions = []Transaction{}
	genesisBlock.PrevHash = ""
	genesisBlock.Hash = calculateHash(genesisBlock)

	blockchain := []Block{genesisBlock}

	transactions := []Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 0.5},
		{Sender: "Bob", Receiver: "Charlie", Amount: 0.2},
		{Sender: "Alice", Receiver: "Dave", Amount: 0.7},
	}

	block := generateBlock(blockchain[len(blockchain)-1], transactions)
	blockchain = append(blockchain, block)

	for _, block := range blockchain {
		fmt.Printf("Bloco #%d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Println("Transações:")
		for _, transaction := range block.Transactions {
			fmt.Printf("- Remetente: %s, Destinatário: %s, Valor: %.2f\n", transaction.Sender, transaction.Receiver, transaction.Amount)
		}
		fmt.Printf("Hash anterior: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("------------------")
	}

	fmt.Printf("O blockchain é válido? %t\n", isChainValid(blockchain))
}
