package help

import (
	"fmt"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
)

type HelpCommand struct {
	CommandService command.CommandService
}

func (h *HelpCommand) Execute(_ []string) error {
	fmt.Println("Available commands:")

	availableCommands := h.CommandService.GetAvailableCommands()
	for cmdName, cmd := range availableCommands {
		fmt.Printf("  %s - %s\n", cmdName, cmd.Description())
	}

	return nil
}

func (h *HelpCommand) Description() string {
	return "Show available commands"
}
