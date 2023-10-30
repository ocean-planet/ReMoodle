package core

import (
	"fmt"

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
