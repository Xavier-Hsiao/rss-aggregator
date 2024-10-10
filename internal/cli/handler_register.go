package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/config"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
	"github.com/google/uuid"
)

func HandlerRegister(s *app.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("expect one argument as username")
	}

	newUser, err := s.DB.CreateUser(
		context.Background(),
		datbase.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.Args[0],
		},
	)
	if err != nil {
		return err
	}

	// Exit with code 1 if a user with the same name already exists
	if newUser.Name == s.Config.CurrentUserName {
		log.Fatal("User already exists")
	}

	// Write the new user to config file
	s.Config.CurrentUserName = newUser.Name
	err = config.Write(*s.Config)
	if err != nil {
		return err
	}

	fmt.Println("User created!")
	fmt.Printf(
		"username: %v\n user id: %v\n created at: %v\n",
		newUser.Name,
		newUser.ID,
		newUser.CreatedAt,
	)

	return nil
}
