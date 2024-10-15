package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
	"github.com/google/uuid"
)

func HandlerFollow(s *app.State, cmd Command) error {
	ctx := context.Background()
	url := cmd.Args[0]

	user, err := s.DB.GetUserByName(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user by name: %v", err)
	}

	feed, err := s.DB.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %v", err)
	}

	feedFollowRow, err := s.DB.CreateFeedFollow(
		ctx,
		datbase.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating feedFollow row: %v", err)
	}

	fmt.Printf("Hey %s, you just follow: %s\n",
		feedFollowRow.UserName,
		feedFollowRow.FeedName,
	)

	return nil
}
