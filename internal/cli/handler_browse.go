package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
)

func HandlerBrowse(s *app.State, cmd Command, user datbase.User) error {
	// Set the default limit to 2
	limit := 2
	if len(cmd.Args) == 1 {
		specifiedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit provided: %v", err)
		}

		limit = specifiedLimit
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), datbase.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to fetch posts for %s: %v", user.Name, err)
	}

	fmt.Printf("Found %d posts for %s\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Println("=====================================")
		fmt.Printf("---- %s ----\n", post.Title)
		fmt.Printf("     %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("Posted on: %s\n", post.PublishedAt)
		fmt.Println("=====================================")
	}

	return nil
}
