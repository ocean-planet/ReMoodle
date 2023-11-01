package login

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

type LoginCommand struct {
	CommandService command.CommandService
}

func (h *LoginCommand) Execute(_ []string) error {

	loadToken, loadTokenErr := core.LoadToken()

	if loadTokenErr == nil || len(loadToken) > 5 {
		fmt.Printf("You are already logged in!\nYour token: %s", loadToken)
		return loadTokenErr
	}
	fmt.Println("Enter your token:")

	reader := bufio.NewReader(os.Stdin)
	token, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	token = strings.TrimSpace(token)

	moodleRepository := moodle.NewMoodleRepository("https://moodle.astanait.edu.kz/webservice/rest/server.php")
	moodleService := moodle.NewMoodleService(moodleRepository)

	userInfo, tokenErr := moodleService.GetUserInfo(token)

	if tokenErr != nil {
		fmt.Println("Token is invalid!")
		return tokenErr
	}
	err = core.SaveToken(token)

	if err != nil {
		return err
	}

	fmt.Printf("Hello %s 👋, your token was saved successfully!", userInfo.FullName)

	return nil
}

func (h *LoginCommand) Description() string {
	return "Login to Moodle"
}
