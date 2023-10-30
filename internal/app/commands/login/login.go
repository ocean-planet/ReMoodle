package login

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
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

	err = SaveToken(token)
	if err != nil {
		return err
	}

	fmt.Println("Token saved successfully")

	return nil
}

func (h *LoginCommand) Description() string {
	return "Login to Moodle"
}

func SaveToken(token string) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	tokenFile := filepath.Join(configDir, "remoodle", "token")
	err = os.MkdirAll(filepath.Dir(tokenFile), 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(tokenFile, []byte(token), 0600)
}

func LoadToken() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	println(configDir)

	tokenFile := filepath.Join(configDir, "remoodle", "token")
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
