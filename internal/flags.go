package internal

import "github.com/urfave/cli/v2"

// Const values
const (
	Debug      string = "debug"
	File              = "File"
	Repository        = "repository"
	Quiet             = "quiet"
	Owner             = "owner"
)

// CreateRepoFlags create owner/repo struct based on file input.
func CreateRepoFlags(owner string) []cli.Flag {
	required := owner == ""
	return []cli.Flag{
		&cli.StringFlag{
			Name:     Owner,
			Required: required,
			Aliases:  []string{"o"},
			Usage:    "Source repository owner",
			Value:    owner,
		},
		&cli.StringFlag{
			Name:     Repository,
			Required: true,
			Aliases:  []string{"r"},
			Usage:    "Repository name",
		},
	}
}
