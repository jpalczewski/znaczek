package label

import "github.com/google/go-github/v31/github"

type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

func NewLabel(l *github.Label) *Label {
	return &Label{*l.Name, *l.Description, *l.Color}
}
