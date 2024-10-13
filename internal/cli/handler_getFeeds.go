package cli

import (
	"context"
	"fmt"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
)

func HandlerGetFeeds(s *app.State, cmd Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting all feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Println("----------")
		fmt.Printf(
			"* Feed Name: %s\n* Feed URL: %s\n* Created By: %s\n",
			feed.Name,
			feed.Url,
			feed.CreatedBy,
		)
	}

	return nil
}
