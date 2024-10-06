package cli

import (
	"errors"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*app.State, Command) error
}

func (c *Commands) Register(name string, f func(s *app.State, cmd Command) error) {
	c.Handlers[name] = f
}

// Execute a given command with state if it exists
func (c *Commands) Run(s *app.State, cmd Command) error {
	handler, exists := c.Handlers[cmd.Name]
	if !exists {
		return errors.New("error: command not found")
	}

	return handler(s, cmd)
}
