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
	fmt.Println("Enter your token:")

	_, tokenExistsErr := core.LoadToken()

	if tokenExistsErr != nil {
		fmt.Println("You are already logged in!")
	}

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
