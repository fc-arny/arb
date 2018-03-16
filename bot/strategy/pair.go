package strategy

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../common"
)

// Pair stratery of two ex and two coins
func Pair(cfg *common.Config) {
	opts := cfg.Strategy.Options
	cycles, _ := strconv.ParseInt(opts["cycles"], 10, 32)

	for i := 0; i < int(cycles); i++ {
		for _, ex := range strings.Split(opts["exchanges"], ",") {
			switch ex {
			case "bittrex":
				fmt.Print("bittrex - OK")
			case "binance":
				fmt.Print("binance - OK")
			default:
				fmt.Print("\n")
				log.Fatalf("Wrong Exchange %q", ex)
			}

		}
	}
	fmt.Print("\n")
}
