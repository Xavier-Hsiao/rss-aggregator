package main

import (
	"fmt"
	"log"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/config"
)

func main() {
	// Initial config file read
	c, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config file")
	}

	currentUserName := "Xavier"

	c.SetUser(currentUserName)

	c, err = config.Read()
	if err != nil {
		log.Fatalf("failed to read config file")
	}

	fmt.Printf("Config content: %v, %v\n", c.DbURL, c.CurrentUserName)
}
