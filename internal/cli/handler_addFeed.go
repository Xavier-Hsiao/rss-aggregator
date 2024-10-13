package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *app.State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("error not enough arguments, expected 2")
	}

	// Get the current user name for fk (user_id) in feeds table
	currentUserName := s.Config.CurrentUserName
	user, err := s.DB.GetUserByName(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	userID := user.ID

	newFeed, err := s.DB.CreateFeed(
		context.Background(),
		datbase.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.Args[0],
			Url:       cmd.Args[1],
			UserID:    userID,
		},
	)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Println("Feed created!")
	fmt.Printf(
		"name: %v\n feed id: %v\n url: %v\n user: %v\n created at: %v\n",
		newFeed.Name,
		newFeed.ID,
		newFeed.Url,
		newFeed.UserID,
		newFeed.CreatedAt,
	)

	return nil
}
