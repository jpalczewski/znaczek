package internal

// Config stores settings stored in a file.
type Config struct {
	Defaults struct {
		Owner string `fig:"owner"`
	}
}
