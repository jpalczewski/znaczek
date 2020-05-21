package internal

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v31/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

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
