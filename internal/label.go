package internal

//Label represents a GH label in simplified manner
type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}
