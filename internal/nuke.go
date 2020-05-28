package internal

import (
	"github.com/urfave/cli/v2"
)

//Nuke removes all labels from repository
func Nuke(c *cli.Context) error {
	ctx, client := getClient()

	labels, _, _ := client.Issues.ListLabels(ctx, c.String(Owner), c.String(Repository), nil)

	for _, x := range labels {
		client.Issues.DeleteLabel(ctx, c.String(Owner), c.String(Repository), *x.Name)

	}

	return nil
}
