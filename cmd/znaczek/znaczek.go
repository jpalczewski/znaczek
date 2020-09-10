package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/jpalczewski/znaczek/internal"
	"github.com/kkyr/fig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {

	var cfg internal.Config
	var debugLevel bool
	path, err := xdg.ConfigFile("znaczek/settings.yaml")
	dir, file := filepath.Split(path)

	if err != nil {
		log.Println(err, path)
	}

	err = fig.Load(&cfg, fig.File(file), fig.Dirs(dir))
	if err != nil {
		log.Println(err)
	}

	repoDetailsFlags := internal.CreateRepoFlags(cfg.Defaults.Owner)

	repoAndFileDetailsFlags := append(repoDetailsFlags, &cli.StringFlag{
		Name:    internal.File,
		Aliases: []string{"f"},
		Value:   "labels.json",
		Usage:   "Related file name",
	})

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "export",
				Aliases: []string{"e"},
				Usage:   "export labels to file",
				Action:  internal.Export,
				Flags:   repoAndFileDetailsFlags,
			},
			{
				Name:    "nuke",
				Action:  internal.Nuke,
				Usage:   "Delete all label in existing repo",
				Aliases: []string{"n"},
				Flags:   repoDetailsFlags,
			},
			{
				Name:    "apply",
				Action:  internal.Apply,
				Usage:   "Apply json file to repo",
				Aliases: []string{"a"},
				Flags:   repoAndFileDetailsFlags,
			},
		},
		Usage: "Managing github labels have never been easier!",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        internal.Debug,
				Aliases:     []string{"d", "v"},
				Value:       false,
				Usage:       "Show further info",
				Destination: &debugLevel,
			},
			&cli.BoolFlag{
				Name:    internal.Quiet,
				Aliases: []string{"q"},
				Value:   false,
				Usage:   "Hide unnecesary info",
			},
		},
	}
	err = app.Run(os.Args)

	logger := internal.CreateLogger(debugLevel)
	zap.ReplaceGlobals(logger.Desugar())
	if err != nil {
		log.Fatal(err)
	}
}
