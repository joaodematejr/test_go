package main

import (
	"fmt"
	"time"
)

type NFT struct {
	ID       uint64
	Owner    string
	Metadata string
}

type NFTContract struct {
	Name   string
	Symbol string
	Owner  string
	Tokens []NFT
}

func createNFT(owner, metadata string) NFT {
	nft := NFT{
		ID:       uint64(time.Now().UnixNano()),
		Owner:    owner,
		Metadata: metadata,
	}
	return nft
}

func (contract *NFTContract) addNFT(owner, metadata string) {
	nft := createNFT(owner, metadata)
	contract.Tokens = append(contract.Tokens, nft)
}

func (contract *NFTContract) totalSupply() uint64 {
	return uint64(len(contract.Tokens))
}

func (contract *NFTContract) getNFTByID(id uint64) (NFT, error) {
	for _, nft := range contract.Tokens {
		if nft.ID == id {
			return nft, nil
		}
	}
	return NFT{}, fmt.Errorf("NFT with ID %d not found", id)
}

func main() {
	contract := NFTContract{
		Name:   "MyNFT",
		Symbol: "NFT",
		Owner:  "Alice",
		Tokens: []NFT{},
	}

	contract.addNFT("Alice", "Metadata 1")
	contract.addNFT("Bob", "Metadata 2")
	contract.addNFT("Charlie", "Metadata 3")

	totalSupply := contract.totalSupply()
	fmt.Println("Total Supply:", totalSupply)

	nftID := uint64(2)
	nft, err := contract.getNFTByID(nftID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("NFT ID:", nft.ID)
		fmt.Println("Owner:", nft.Owner)
		fmt.Println("Metadata:", nft.Metadata)
	}
}
