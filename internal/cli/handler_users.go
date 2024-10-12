package cli

import (
	"context"
	"fmt"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
)

func HanlderUsers(s *app.State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting all users from database: %v", err)
	}

	currentUser := s.Config.CurrentUserName
	for _, user := range users {
		// Check if the user is currentUser
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", currentUser)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
