package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ocean-planet/ReMoodle/internal/app/logger"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
)

type App struct {
	CommandService command.Service
}

func NewApp(service command.Service) *App {
	return &App{
		CommandService: service,
	}
}

const (
	ConfigDirName = "remoodle"
	TokenFileName = "user_token"
)

func (a *App) RegisterCommand(name string, cmd command.Command) {
	const op = "core.RegisterCommand"
	a.CommandService.RegisterCommand(name, cmd)
	logger.LogWithPrefix(op, "Command ", name, " registered")
}

func (a *App) Run(args []string) error {
	const op = "core.Run"
	if len(args) < 1 {
		logger.LogWithPrefix(op, "Application ran with 0 arguments. Help page was shown")
		fmt.Println("Usage: remoodle <command> [args]")
		cmd, _ := a.CommandService.GetCommand("help")
		err := cmd.Execute(nil)
		if err != nil {
			return err
		}
		return nil
	}

	commandName := args[0]
	logger.LogWithPrefix(op, "Got command:", commandName)

	cmd, found := a.CommandService.GetCommand(commandName)
	if !found {
		logger.LogWithPrefix(op, "Command:", commandName, "is unknown")
		fmt.Println("Unknown command:", commandName)
		return nil
	}

	err := cmd.Execute(args[1:])
	if err != nil {
		logger.LogWithPrefix(op, "Error while executing command:", err.Error())
		fmt.Println("Error executing command:", err)
		return err
	}

	return nil
}

func SaveToken(token string) error {
	const op = "core.SaveToken"
	configDir, err := os.UserConfigDir()
	if err != nil {
		logger.LogWithPrefix(op, "Failed to read configDir:", err.Error())
		return err
	}

	tokenFile := filepath.Join(configDir, ConfigDirName, TokenFileName)
	err = os.MkdirAll(filepath.Dir(tokenFile), 0755)
	if err != nil {
		logger.LogWithPrefix(op, "Failed to create tokenFile:", err.Error())
		return err
	}

	return os.WriteFile(tokenFile, []byte(token), 0600)
}

func LoadToken() (string, error) {
	const op = "core.LoadToken"
	configDir, err := os.UserConfigDir()
	if err != nil {
		logger.LogWithPrefix(op, "Failed to read configDir:", err.Error())
		return "", err
	}

	tokenFile := filepath.Join(configDir, ConfigDirName, TokenFileName)
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		logger.LogWithPrefix(op, "Failed to read tokenFile:", err.Error())
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func DeleteToken() error {
	const op = "core.DeleteToken"
	configDir, err := os.UserConfigDir()
	if err != nil {
		logger.LogWithPrefix(op, "Failed to read configDir:", err.Error())
		return err
	}

	tokenFile := filepath.Join(configDir, ConfigDirName, TokenFileName)
	err = os.Remove(tokenFile)
	if err != nil {
		logger.LogWithPrefix(op, "Failed to delete tokenfile:", err.Error())
		return err
	}

	return nil
}
