package help

import (
	"fmt"

	"os"
	"text/tabwriter"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
)

type Command struct {
	CommandService command.Service
}

func (c *Command) Execute(_ []string) error {
	fmt.Println("Available commands:")

	availableCommands := c.CommandService.GetAvailableCommands()
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	for cmdName, cmd := range availableCommands {
		row := fmt.Sprintf("    %s \t  %s", cmdName, cmd.Description())
		if _, err := fmt.Fprintln(tw, row); err != nil {
			return fmt.Errorf("error writing to tabwriter: %v", err)
		}
	}
	if err := tw.Flush(); err != nil {
		return fmt.Errorf("error flushing tabwriter: %v", err)
	}

	return nil
}

func (c *Command) Description() string {
	return "Show available commands"
}
