package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
)

type App struct {
	CommandService command.CommandService
}

func NewApp(service command.CommandService) *App {
	return &App{
		CommandService: service,
	}
}

const (
	ConfigDirName = "remoodle"
	TokenFileName = "user_token"
)

func (a *App) RegisterCommand(name string, cmd command.Command) {
	a.CommandService.RegisterCommand(name, cmd)
}

func (a *App) Run(args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: myapp <command> [args]")
		return nil
	}

	commandName := args[0]

	cmd, found := a.CommandService.GetCommand(commandName)
	if !found {
		fmt.Println("Unknown command:", commandName)
		return nil
	}

	err := cmd.Execute(args[1:])
	if err != nil {
		fmt.Println("Error executing command:", err)
		return err
	}

	return nil
}

func SaveToken(token string) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	tokenFile := filepath.Join(configDir, ConfigDirName, TokenFileName)
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

	tokenFile := filepath.Join(configDir, ConfigDirName, TokenFileName)
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
