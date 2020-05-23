package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/google/go-github/v31/github"
	"github.com/urfave/cli/v2"
)

func Apply(c *cli.Context) error {
	ctx, client := getClient()

	rawFile, error := ioutil.ReadFile(c.String("file"))

	if error != nil {
		log.Fatal(error)
	}

	var labels []Label
	err := json.Unmarshal(rawFile, &labels)

	if err != nil {
		log.Fatal(err)
	}

	for _, x := range labels {
		label := github.Label{Name: &x.Name, Description: &x.Description, Color: &x.Color}
		client.Issues.CreateLabel(ctx, c.String("owner"), c.String("repository"), &label)

	}

	return nil
}
