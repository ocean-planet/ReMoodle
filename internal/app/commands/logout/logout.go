package logout

import (
	"fmt"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
)

type LogoutCommand struct {
	CommandService command.CommandService
}

func (l *LogoutCommand) Description() string {
	return "delete your token"
}

func (l *LogoutCommand) Execute(_ []string) error {
	loadToken, loadTokenErr := core.LoadToken()

	if loadTokenErr != nil || len(loadToken) > 5 {
		deleteTokenErr := core.DeleteToken()

		if deleteTokenErr != nil {
			fmt.Println("Something went wrong!")
			return deleteTokenErr
		}

		fmt.Printf("Your token %s has been removed!", loadToken)
		return nil
	} else {
		fmt.Printf("You was not logged in!")
		return nil
	}
}
