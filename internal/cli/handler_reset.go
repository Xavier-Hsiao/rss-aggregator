package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/config"
)

func HandlerReset(s *app.State, cmd Command) error {
	// Delete users from database
	err := s.DB.DeleteUsers(context.Background())
	if err != nil {
		log.Fatal("Failed to delete users")
	}

	// Delete the current user record in config json file
	s.Config.CurrentUserName = ""

	err = config.Write(*s.Config)
	if err != nil {
		return fmt.Errorf("error unapdating config file")
	}

	return nil
}
