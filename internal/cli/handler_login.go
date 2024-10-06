package cli

import (
	"errors"
	"fmt"

	"github.com/Xavier-Hsiao/rss-aggregator/internal/app"
)

func HandlerLogin(s *app.State, cmd Command) error {
	if len(cmd.args) == 0 {
		return errors.New("Error: expect one argument")
	}

	currentUserName := cmd.args[0]
	s.Config.SetUser(currentUserName)

	fmt.Println("The user name has been set.")

	return nil
}
