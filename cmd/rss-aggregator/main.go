package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/cli"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/config"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
	_ "github.com/lib/pq"
)

func main() {
	// Initial config file read
	c, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config file:\n", err)
	}

	// Connect to database
	db, err := sql.Open("postgres", c.DbURL)
	if err != nil {
		log.Fatalf("failed to connect to database:\n", err)
	}

	dbQueries := datbase.New(db)

	state := &app.State{
		Config: &c,
		DB:     dbQueries,
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
