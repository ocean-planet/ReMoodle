package login

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
)

type LoginCommand struct {
	CommandService command.CommandService
}

func (h *LoginCommand) Execute(_ []string) error {
	fmt.Println("Enter your token:")

	reader := bufio.NewReader(os.Stdin)
	token, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	token = strings.TrimSpace(token)

	err = core.SaveToken(token)
	if err != nil {
		return err
	}

	fmt.Println("Token saved successfully")

	return nil
}

func (h *LoginCommand) Description() string {
	return "Login to Moodle"
}
