package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
)

func HandlerLogin(s *app.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("expect one argument as username")
	}

	currentUserName := cmd.Args[0]

	// Check if the given user exists in database
	_, err := s.DB.GetUserByName(context.Background(), currentUserName)
	if err != nil {
		log.Fatalf("couldn't find %v\n", currentUserName)
	}

	err = s.Config.SetUser(currentUserName)
	if err != nil {
		return fmt.Errorf("failed to set current user: %v", err)
	}

	fmt.Println("User has been switched!")

	return nil
}
