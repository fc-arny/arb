package exchange

// Interface common interface for any Exchange
type Interface interface {
	GetName() string
	GetOrderBook(market string) (OrderBook, error)
	GetBalance(currency string) Balance
	// CreateOrder() error
	// GetBalance()
	// GetFee
	// GetWalletStatus
	// Cancel
	// GetOpenOrders
	// OrderStatus
	// GetMarket
}

type OrderBook struct {
	Bids []Order `json:"bids"`
	Asks []Order `json:"asks"`
}

type Order struct {
	Price    float64 `json:",string"`
	Quantity float64 `json:",string"`
}

// Balance of account
type Balance struct {
	Currency string  `json:"Currency"`
	Amount   float64 `json:"Amount"`
}

// const Items = []{BinanceAdapter, BittexAdapter}
