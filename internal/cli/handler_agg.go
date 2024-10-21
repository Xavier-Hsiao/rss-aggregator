package cli

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func HandlerAgg(s *app.State, cmd Command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}

	log.Printf("Collecting feeds every %s...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *app.State) {
	ctx := context.Background()
	// Get the next feed to fetch from DB
	feed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Println("Failed to get next feeds")
		return
	}
	log.Println("Found a feed to fetch!")

	// Mark the feed as fetched
	_, err = s.DB.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		log.Printf("Failed to mark feed %v as fetched\n", feed.Name)
		return
	}

	// Fetch the feed using URL
	feedData, err := app.FetchFeed(ctx, feed.Url)
	if err != nil {
		log.Printf("Failed to fetch feed %v\n", feed.Name)
		return
	}

	// Iterate over the items in the feed and save posts to database
	for _, item := range feedData.Channel.Item {
		// Convert published_at to time
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Failed to convert published_at time: %v", err)
		}

		_, err = s.DB.CreatePost(ctx, datbase.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			// Skip duplicate URL error
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue
				}
			}
			log.Printf("Failed to create post: %v", err)
			continue
		}
	}
	fmt.Printf("Feed %s collected: %d posts found!\n", feed.Name, len(feedData.Channel.Item))
}
