package main

import (
	"log"
	"os"

	"github.com/jpalczewski/znaczek/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "export",
				Aliases: []string{"e"},
				Usage:   "export labels to file",
				Action:  internal.Export,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "owner",
						Required: true,
						Aliases:  []string{"o"},
						Usage:    "Source repository owner",
					},
					&cli.StringFlag{
						Name:     "repository",
						Required: true,
						Aliases:  []string{"r"},
						Usage:    "Repository name",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"f"},
						Value:   "labels.json",
						Usage:   "Output file name",
					},
				},
			},
			{
				Name:    "nuke",
				Action:  internal.Nuke,
				Usage:   "Delete all label in existing repo",
				Aliases: []string{"n"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "owner",
						Required: true,
						Aliases:  []string{"o"},
						Usage:    "Source repository owner",
					},
					&cli.StringFlag{
						Name:     "repository",
						Required: true,
						Aliases:  []string{"r"},
						Usage:    "Repository name",
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
				Usage:   "Hide unnecesary info",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
