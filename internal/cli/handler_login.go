package cli

import (
	"fmt"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
)

func HandlerLogin(s *app.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("expect one argument as username")
	}

	currentUserName := cmd.Args[0]
	err := s.Config.SetUser(currentUserName)
	if err != nil {
		return fmt.Errorf("failed to set current user: %v", err)
	}

	fmt.Println("User has been switched!")

	return nil
}
