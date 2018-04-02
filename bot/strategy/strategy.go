package strategy

import (
	"fmt"

	"../common"
	"../exchange"
)

// Strategy interface
type Strategy struct {
	Config    common.Config
	Exchanges map[string]exchange.Interface
}

// Init initialize
func (s *Strategy) Init() *Strategy {
	fmt.Println("Init:")
	s.setupExchanges()

	return s
}

// Options of strategy
func (s *Strategy) Options() map[string]string {
	return s.Config.Strategy.Options
}

// Type of strategy
func (s *Strategy) Type() string {
	return s.Config.Strategy.Type
}

// Run strategy
func (s *Strategy) Run() {
	fmt.Println("#### STRAGEGY #### ")
	switch s.Type() {
	case "pair":

		p := PairStrategy{Strategy: *s}
		p.Start()
	}
	fmt.Println("#### STRAGEGY stop #### ")
}

func (s *Strategy) setupExchanges() {
	s.Exchanges = make(map[string]exchange.Interface)

	fmt.Println("- setup exchanges")
	for key, config := range s.Config.Auth {
		switch key {
		case "bittrex":
			fmt.Println("-- bittrex")
			s.Exchanges[key] = exchange.BittexAdapter{Config: config}
		case "binance":
			fmt.Println("-- binance")
			s.Exchanges[key] = exchange.BinanceAdapter{Config: config}
		default:
			panic("Wrong exchange - " + key)
		}
	}
}
