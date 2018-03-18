package strategy

import (
	"fmt"
	// "log"
	"strconv"
	// "strings"

	"../common"
	"../exchange"
)

type PairStrategy struct {
	Strategy
}

// func (p *PairStrategy) setup(cfg *common.Config) {
// 	opts := cfg.Strategy.Options
// 	cycles, _ := strconv.ParseInt(opts["cycles"], 10, 32)
// 	for _, ex := range strings.Split(opts["exchanges"], ",") {

// 			trader.BittexAdapter.GetOrderBook()
// 			trader.BinanceAdapter.GetOrderBook()

// 	}
// }

// Pair stratery of two ex and two coins
func Pair(cfg *common.Config) {
	opts := cfg.Strategy.Options
	cycles, _ := strconv.ParseInt(opts["cycles"], 10, 32)

	for i := 0; i < int(cycles); i++ {
		fmt.Print(".")
		p := exchange.BittexAdapter{}
		p.GetOrderBook()
		// exchange.BittexAdapter.GetOrderBook()
		// trader.BinanceAdapter.GetOrderBook()
	}
	fmt.Print("\n")
}
