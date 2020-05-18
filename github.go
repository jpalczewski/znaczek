package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/kyokomi/emoji"
	. "github.com/logrusorgru/aurora"
)

func export(c *cli.Context) error {
	fmt.Println("I should export something.")
	if c.Bool("debug") {
		emoji.Println(Yellow("But not yet :cowboy_hat_face:"))
	}

	return nil
}
