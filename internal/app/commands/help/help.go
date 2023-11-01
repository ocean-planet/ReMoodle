package help

import (
	"fmt"

	"os"
	"text/tabwriter"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
)

type HelpCommand struct {
	CommandService command.CommandService
}

func (h *HelpCommand) Execute(_ []string) error {
	fmt.Println("Available commands:")

	availableCommands := h.CommandService.GetAvailableCommands()
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	for cmdName, cmd := range availableCommands {
		row := fmt.Sprintf("    %s \t  %s", cmdName, cmd.Description())
		fmt.Fprintln(tw, row)
	}
	tw.Flush()

	return nil
}

func (h *HelpCommand) Description() string {
	return "Show available commands"
}
