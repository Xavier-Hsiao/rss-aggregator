package cli

import (
	"context"
	"log"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
)

func HandlerDeleteFeedFollowsForUser(s *app.State, cmd Command, user datbase.User) error {
	url := cmd.Args[0]

	err := s.DB.DeleteFollowsForUser(context.Background(), datbase.DeleteFollowsForUserParams{
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		log.Fatal("Failed to delete users")
	}

	return nil
}
