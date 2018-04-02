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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				DefaultText: "./config.yml",
				Usage:       "Load configuration from `FILE.yml`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run bot with your config",
				Action: func(c *cli.Context) error {
					config.LoadFrom(c.String("config"))

					cstrategy := strategy.Strategy{Config: config}
					cstrategy.Init().Run()

					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
