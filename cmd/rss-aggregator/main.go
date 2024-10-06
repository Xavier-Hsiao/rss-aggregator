package main

import (
	"log"
	"os"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/cli"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/config"
)

func main() {
	// Initial config file read
	c, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config file")
	}

	state := &app.State{
		Config: &c,
	}

	commands := &cli.Commands{
		Handlers: make(map[string]func(*app.State, cli.Command) error),
	}
	commands.Register("login", cli.HandlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough argument provided")
	}

	command := &cli.Command{
		Name: args[1],
		Args: args[2:],
	}
	err = commands.Run(state, *command)
	if err != nil {
		log.Fatal(err)
	}
}
