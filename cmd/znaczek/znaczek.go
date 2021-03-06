package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/jpalczewski/znaczek/internal"
	"github.com/kkyr/fig"
	"github.com/urfave/cli/v2"
)

func main() {

	var cfg internal.Config

	path, err := xdg.ConfigFile("znaczek/settings.yaml")
	dir, file := filepath.Split(path)

	if err != nil {
		log.Fatal(err, path)
	}

	err = fig.Load(&cfg, fig.File(file), fig.Dirs(dir))
	if err != nil {
		log.Fatalln(err)
	}

	repoDetailsFlags := internal.CreateRepoFlags(cfg.Defaults.Owner)

	repoAndFileDetailsFlags := append(repoDetailsFlags, &cli.StringFlag{
		Name:    "file",
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
				Name:    internal.Debug,
				Aliases: []string{"d", "v"},
				Value:   false,
				Usage:   "Show further info",
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
	if err != nil {
		log.Fatal(err)
	}
}
