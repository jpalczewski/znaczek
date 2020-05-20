package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v31/github"
	"github.com/joho/godotenv"
	"github.com/jpalczewski/znaczek/label"
	"github.com/kyokomi/emoji"
	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

func export(c *cli.Context) error {
	fmt.Println("I should export something.")
	if c.Bool("debug") {
		emoji.Println(Yellow("But not yet :cowboy_hat_face:"))
	}
	ctx, client := getClient()

	sourceLabels, _, err := client.Issues.ListLabels(ctx, c.String("owner"), c.String("repository"), nil)
	if err != nil {
		log.Fatal(err)
	}

	exportLabels := make([]label.Label, len(sourceLabels))
	for i, v := range sourceLabels {
		exportLabels[i] = *label.NewLabel(v)
	}
	out, _ := json.MarshalIndent(exportLabels, "  ", "    ")

	if !c.Bool("quiet") {
		fmt.Println(string(out))
	}

	handle, _ := os.Create(c.String("output"))
	handle.WriteString(string(out))
	// for i, x := range ret {
	// 	lab := *label.NewLabel(x)
	// 	test, err := json.Marshal(lab)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(i, string(test))
	// }
	return nil
}

func getClient() (context.Context, *github.Client) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	githubToken := os.Getenv("GH_TOKEN")

	fmt.Println("Test", githubToken)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return ctx, client
}
