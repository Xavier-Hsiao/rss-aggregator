package cli

import "github.com/Xavier-Hsiao/rss-aggregator/internal/app"

type Command struct {
	name string
	args []string
}

type Commands struct {
	handlers map[string]func(*app.State, Command) error
}
