package exchange

import (
	"fmt"
	"strings"

	"github.com/pdepip/go-binance/binance"
)

// BinanceAdapter adap`ter
type BinanceAdapter struct {
	Config map[string]string
}

// GetName echange name
func (b BinanceAdapter) GetName() string {
	return "Binance"
}

// GetFee trading fee for transaction
func (b BinanceAdapter) GetFee() float64 {
	return 0.1 / 100
}

// GetOrderBook get
func (b BinanceAdapter) GetOrderBook(market string) (orderBook OrderBook, err error) {
	query := binance.OrderBookQuery{
		Symbol: b.aMarket(market),
		Limit:  100,
	}

	markets, err := b.client().GetOrderBook(query)
	for _, buy := range markets.Bids {
		orderBook.Bids = append(orderBook.Bids, Order{Price: buy.Price, Quantity: buy.Quantity})
	}

	for _, ask := range markets.Asks {
		orderBook.Asks = append(orderBook.Asks, Order{Price: ask.Price, Quantity: ask.Quantity})
	}

	return
}

func (b BinanceAdapter) GetBalance(currency string) Balance {
	accountUnfo, _ := b.client().GetAccountInfo()

	ammount := 0.0
	for _, balance := range accountUnfo.Balances {
		if balance.Asset == currency {
			ammount = balance.Free
			break
		}

	}
	fmt.Println("BinanceAdapter: GetBalance - ", ammount, currency)

	return Balance{currency, ammount}
}

func (b BinanceAdapter) client() *binance.Binance {
	return binance.New(b.Config["secret"], b.Config["key"])
}

func (b BinanceAdapter) aMarket(market string) string {
	return strings.Join(strings.Split(market, "-"), "")
}
