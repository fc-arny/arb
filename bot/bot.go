package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"

	"./common"
	"./strategy"
)

// "github.com/pdepip/go-binance/binance"

func main() {
	var config common.Config

	app := &cli.App{
		Name:    "bot",
		Usage:   "CryptoBOT",
		Version: "0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "fc_arny",
				Email: "arthur.shcheglov@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE.yml`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run bot with your config",
				Action: func(c *cli.Context) error {
					config.LoadFrom(c.String("config"))

					switch config.Strategy.Type {
					case "pair":
						fmt.Println("Start PAIR strategy")
						strategy.Pair(&config)
						fmt.Println("End PAIR strategy")
					}

					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}

// Bittrex client
// bittrex := bittrex.New(API_KEY, API_SECRET)

// pair := "BTC-ETH"
// // os.Getenv("PAIR") ||

// fmt.Println("BITREX START ==========================")
// bOrderBook, err := bittrex.GetOrderBook(pair, "both")
// fmt.Println(err, bOrderBook)

// fmt.Println("BITREX END ==========================")

// fmt.Println("EXMO START ==========================")
// eOrderBook, err := exmo.Request("order_book", exmo.APIParams{"pair": "ETH_BTC"})
// fmt.Println(err, eOrderBook)
// fmt.Println("EXMO END ==========================")
