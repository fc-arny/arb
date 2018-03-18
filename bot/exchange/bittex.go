package exchange

import (
	"fmt"

	"github.com/toorop/go-bittrex"
)

const (
	apiKey    = "a58a682e0ad64616a758899e0dac0d38"
	apiSecret = "6edfc1db8452494ea74dade77abc6931"
)

// BittexAdapter adapter
type BittexAdapter struct {
	Options map[string]string
	lib     bittrex.Bittrex
}

func (b BittexAdapter) New() {
	fmt.Println("New bittrix")
}

// GetOrderBook order book
func (b BittexAdapter) GetOrderBook() {
	fmt.Println("Bittrex: GetOrder")

	// fmt.Println("Bittrex - orders")
	// // Bittrex client
	// client := client.New(apiKey, apiSecret)

	// // Get markets
	// markets, err := client.GetMarkets()
	// fmt.Println(err, markets)
}
