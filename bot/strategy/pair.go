package strategy

import (
	"fmt"
	"strconv"
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
		orderBooks := make(map[string]exchange.OrderBook)

		for _, cEx := range exchanges {
			orderBook, err := p.Exchanges[cEx].GetOrderBook(opts["market"])

			orderBooks[cEx] = orderBook

			if err != nil {
				fmt.Print("E")
				time.Sleep(5 * time.Second)
				break
			}
		}

		if len(orderBooks) < 2 {
			break
		}

		p.makeDecision(orderBooks, exchanges[0], exchanges[1])
		p.makeDecision(orderBooks, exchanges[1], exchanges[0])
	}
}

func (p *PairStrategy) makeDecision(orderBooks map[string]exchange.OrderBook, bid string, ask string) {
	r1 := orderBooks[bid].Bids[1].Price * (1.0 - p.Exchanges[bid].GetFee())
	r2 := orderBooks[ask].Asks[1].Price * (1.0 + p.Exchanges[ask].GetFee())
	r := r1/r2 - 1

	fmt.Printf("Buy@%s - %v, Sell@%s - %v, minTradeRelative = %v \n", bid, orderBooks[bid].Bids[1].Price, ask, orderBooks[ask].Asks[1].Price, r)

	if r > 0 {
		opts := p.Options()
		minTradeRelative, _ := strconv.ParseFloat(opts["minTradeRelative"], 64)

		if r > minTradeRelative {
			fmt.Println("Relative - OK:", r)
		}
	}

	// fmt.Print(".")

	// if diff1 > minTradeSize {
	// 	// orderBooks[0].Bids[1].Quantity, orderBooks[1].Asks[1].Quantity
	// 	fmt.Println("BUY ON "+p.Exchanges[1].GetName(), ",", diff1)
	// }

	// if diff2 > minTradeSize {
	// 	// fmt.Println("BUY ON "+exchanges[0], ",", diff2)
	// }
}

func (p *PairStrategy) getCurrency() string {
	opts := p.Options()
	return strings.Split(opts["market"], "-")[1]
}
