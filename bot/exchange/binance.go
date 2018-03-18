package exchange

import (
	"fmt"

	"github.com/pdepip/go-binance/binance"
)

// BinanceAdapter adapter
type BinanceAdapter struct {
	lib binance.Binance
}

// GetOrderBook get orders
func (b BinanceAdapter) GetOrderBook() {
	fmt.Println("Binance: GetOrder")
}