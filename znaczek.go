package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "export",
				Aliases: []string{"e"},
				Usage:   "export labels to file",
				Action:  export,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "owner",
						Required: true,
						Aliases:  []string{"o"},
					},
					&cli.StringFlag{
						Name:     "repository",
						Required: true,
						Aliases:  []string{"r"},
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"f"},
						Value:   "labels.json",
					},
				},
			},
		},
		Usage: "Managing github labels have never been easier!",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d", "v"},
				Value:   false,
				Usage:   "Show further info",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Value:   false,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
