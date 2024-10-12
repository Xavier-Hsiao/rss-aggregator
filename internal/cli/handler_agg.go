package cli

import (
	"context"
	"fmt"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
)

func HandlerAgg(s *app.State, cmd Command) error {
	ctx := context.Background()
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := app.FetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error fetching a feed: %v", err)
	}

	fmt.Printf("Feed: %v\n", *feed)

	return nil
}
