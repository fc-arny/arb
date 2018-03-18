package exchange

// Exchange common interface for any Exchange
type Interface interface {
	GetOrderBook()
	// CreateOrder() error
	// 	. getBalance
	// 3. GetFee
	// 4. GetWalletStatus
	// 5. CreateOrder
	// 6. Cancel
	// 7. GetOpenOrders
	// 8. OrderStatus
	// 9. GetMarket
}

// const Items = []{BinanceAdapter, BittexAdapter}
