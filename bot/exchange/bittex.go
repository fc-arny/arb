package exchange

import (
	"fmt"
	"strings"

	"github.com/toorop/go-bittrex"
)

// BittexAdapter adapter
type BittexAdapter struct {
	Config map[string]string
}

func (b BittexAdapter) GetName() string {
	return "Bittrex"
}

// GetOrderBook order book
func (b BittexAdapter) GetOrderBook(market string) (orderBook OrderBook, err error) {
	markets, err := b.client().GetOrderBook(b.aMarket(market), "both")

	for _, buy := range markets.Buy {
		price, _ := buy.Rate.Float64()
		quantity, _ := buy.Quantity.Float64()

		orderBook.Bids = append(orderBook.Bids, Order{Price: price, Quantity: quantity})
	}

	for _, sell := range markets.Sell {
		price, _ := sell.Rate.Float64()
		quantity, _ := sell.Quantity.Float64()

		orderBook.Asks = append(orderBook.Asks, Order{Price: price, Quantity: quantity})
	}

	return orderBook, err
}

func (b BittexAdapter) GetBalance(currency string) Balance {
	balances, _ := b.client().GetBalance(currency)
	// fmt.Println(balances.Available)
	fmt.Println("BittexAdapter: GetBalance = ", balances.Available, currency)

	amount, _ := balances.Available.Float64()

	return Balance{currency, amount}
}

func (b BittexAdapter) aMarket(market string) string {
	items := strings.Split(market, "-")
	items[0], items[1] = items[1], items[0]
	return strings.Join(items, "-")
}

func (b BittexAdapter) client() *bittrex.Bittrex {
	return bittrex.New(b.Config["key"], b.Config["secret"])
}
