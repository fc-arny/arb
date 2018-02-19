package main

import (
	"fmt"

	"ex/suppliers/exmo"

	"github.com/toorop/go-bittrex"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// Bittrex client
	bittrex := bittrex.New(API_KEY, API_SECRET)

	pair := "BTC-ETH"
	// os.Getenv("PAIR") ||

	fmt.Println("BITREX START ==========================")
	bOrderBook, err := bittrex.GetOrderBook(pair, "both")
	fmt.Println(err, bOrderBook)

	fmt.Println("BITREX END ==========================")

	fmt.Println("EXMO START ==========================")
	eOrderBook, err := exmo.Request("order_book", exmo.APIParams{"pair": "ETH_BTC"})
	fmt.Println(err, eOrderBook)
	fmt.Println("EXMO END ==========================")
}
