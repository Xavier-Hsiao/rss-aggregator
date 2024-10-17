package cli

import (
	"context"
	"fmt"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
)

func HandlerGetFeedFollowsForUser(s *app.State, cmd Command, user datbase.User) error {
	ctx := context.Background()

	FeedFollows, err := s.DB.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows for given user: %v", err)
	}

	fmt.Printf("Hey %s, you're following:\n", s.Config.CurrentUserName)
	for i, feedFollow := range FeedFollows {
		fmt.Printf("%d. %s\n", i+1, feedFollow.FeedName)
	}

	return nil
}
