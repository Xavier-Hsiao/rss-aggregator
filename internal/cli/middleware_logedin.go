package cli

import (
	"context"
	"fmt"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
)

// Higher order function
// Take a handler of "logged in" type -> user
// Return a normal handler
func MiddlewareLoggedIn(
	handler func(s *app.State, cmd Command, user datbase.User) error,
) func(s *app.State, cmd Command) error {
	return func(s *app.State, cmd Command) error {
		currentUserName := s.Config.CurrentUserName
		user, err := s.DB.GetUserByName(context.Background(), currentUserName)
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		return handler(s, cmd, user)
	}
}
