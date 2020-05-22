package internal

import (
	"github.com/urfave/cli/v2"
)

func Nuke(c *cli.Context) error {
	ctx, client := getClient()

	labels, _, _ := client.Issues.ListLabels(ctx, c.String("owner"), c.String("repository"), nil)

	for _, x := range labels {
		client.Issues.DeleteLabel(ctx, c.String("owner"), c.String("repository"), *x.Name)

	}

	return nil
}
