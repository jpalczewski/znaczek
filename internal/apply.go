package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/google/go-github/v31/github"
	"github.com/urfave/cli/v2"
)

// Apply loads content of specified file and applies it to repository.
func Apply(c *cli.Context) error {
	ctx, client := getClient()

	rawFile, error := ioutil.ReadFile(c.String(File))

	if error != nil {
		log.Println("apply")
		log.Fatal(error)
	}

	var labels []Label
	err := json.Unmarshal(rawFile, &labels)

	if err != nil {
		log.Fatal(err)
	}

	for _, x := range labels {
		label := github.Label{Name: &x.Name, Description: &x.Description, Color: &x.Color}
		client.Issues.CreateLabel(ctx, c.String(Owner), c.String(Repository), &label)

	}

	return nil
}
