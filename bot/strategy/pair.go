package strategy

import (
	"fmt"
	"strings"
	"time"

	"../exchange"
)

// PairStrategy is a strategy of two exchanges
type PairStrategy struct {
	Strategy
}

// Start stratery of two ex and two coins
func (p *PairStrategy) Start() {
	opts := p.Options()
	exchanges := strings.Split(opts["exchanges"], ",")

	for _, cEx := range exchanges {
		p.Exchanges[cEx].GetBalance(p.getCurrency())
	}

	for {
		var orderBooks []exchange.OrderBook

		for _, cEx := range exchanges {
			orderBook, err := p.Exchanges[cEx].GetOrderBook(opts["market"])

			orderBooks = append(orderBooks, orderBook)

			if err != nil {
				fmt.Print("E")
				time.Sleep(5 * time.Second)
				break
			}
		}

		if len(orderBooks) < 2 {
			break
		}
		p.makeDecision(orderBooks, exchanges)

	}
}

func (p *PairStrategy) getCurrency() string {
	opts := p.Options()
	return strings.Split(opts["market"], "-")[1]
}

func (p *PairStrategy) makeDecision(orderBooks []exchange.OrderBook, exchanges []string) {
	diff1 := orderBooks[0].Bids[1].Price - orderBooks[1].Asks[1].Price
	diff2 := orderBooks[1].Bids[1].Price - orderBooks[0].Asks[1].Price

	fmt.Print(".")

	if diff1 > 0 {
		// orderBooks[0].Bids[1].Quantity, orderBooks[1].Asks[1].Quantity
		fmt.Println("BUY ON "+exchanges[1], ",", diff1)
	}

	if diff2 > 0 {
		fmt.Println("BUY ON "+exchanges[0], ",", diff2)
	}
}
