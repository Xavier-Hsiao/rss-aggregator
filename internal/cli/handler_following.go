package cli

import (
	"context"
	"fmt"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
)

func HandlerGetFeedFollowsForUser(s *app.State, cmd Command) error {
	ctx := context.Background()
	user, err := s.DB.GetUserByName(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user by name: %v", err)
	}

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
