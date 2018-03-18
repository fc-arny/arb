package main

import (
	// "fmt"
	"os"

	"gopkg.in/urfave/cli.v2"

	"./common"
	"./strategy"
)

// "github.com/pdepip/go-binance/binance"
var config common.Config

// func main() {
// 	config.LoadFrom("./config.yml")

// 	cstrategy := strategy.Strategy{Config: config}
// 	cstrategy.Init()

// }

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
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run bot with your config",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "config",
						Aliases:     []string{"c"},
						DefaultText: "./config.yml",
						Usage:       "Load configuration from `FILE.yml`",
					},
				},
				Action: func(c *cli.Context) error {
					config.LoadFrom(c.String("config"))

					cstrategy := strategy.Strategy{Config: config}
					cstrategy.Init().Run()

					// switch config.Strategy.Type {
					// case "pair":
					// 	fmt.Println("Start PAIR strategy")
					// 	strategy.Pair(&config)
					// 	fmt.Println("End PAIR strategy")
					// }

					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
