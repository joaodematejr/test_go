package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

type Auction struct {
	ID            uint64
	Item          string
	Seller        string
	HighestBid    *big.Int
	HighestBidder string
	EndTime       time.Time
}

type DApp struct {
	Auctions []Auction
}

func (dapp *DApp) createAuction(item string, seller string, endTime time.Time) {
	id := uint64(len(dapp.Auctions) + 1)
	auction := Auction{
		ID:            id,
		Item:          item,
		Seller:        seller,
		HighestBid:    big.NewInt(0),
		HighestBidder: "",
		EndTime:       endTime,
	}
	dapp.Auctions = append(dapp.Auctions, auction)
}

func (dapp *DApp) placeBid(auctionID uint64, bidder string, amount *big.Int) error {
	if auctionID >= uint64(len(dapp.Auctions)) {
		return fmt.Errorf("Auction with ID %d does not exist", auctionID)
	}
	auction := &dapp.Auctions[auctionID]
	if auction.EndTime.Before(time.Now()) {
		return fmt.Errorf("Auction with ID %d has ended", auctionID)
	}
	if amount.Cmp(auction.HighestBid) > 0 {
		auction.HighestBid = amount
		auction.HighestBidder = bidder
	}
	return nil
}

func main() {
	dapp := DApp{}

	dapp.createAuction("Artwork", "Alice", time.Now().Add(time.Hour))

	rand.Seed(time.Now().UnixNano())
	bidder1 := "Bob"
	bidder2 := "Charlie"

	for i := 0; i < 5; i++ {
		amount := big.NewInt(int64(rand.Intn(100) + 1))
		err := dapp.placeBid(0, bidder1, amount)
		if err != nil {
			fmt.Println("Erro ao colocar lance:", err)
		} else {
			fmt.Printf("%s colocou um lance de %s no leilão\n", bidder1, amount.String())
		}
		time.Sleep(1 * time.Second)

		amount = big.NewInt(int64(rand.Intn(100) + 1))
		err = dapp.placeBid(0, bidder2, amount)
		if err != nil {
			fmt.Println("Erro ao colocar lance:", err)
		} else {
			fmt.Printf("%s colocou um lance de %s no leilão\n", bidder2, amount.String())
		}
		time.Sleep(1 * time.Second)
	}

	auction := dapp.Auctions[0]
	fmt.Printf("\nLeilão encerrado para o item '%s'\n", auction.Item)
	fmt.Printf("Vendido para '%s' por %s\n", auction.HighestBidder, auction.HighestBid.String())
}
