package whoami

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

type WhoamiCommand struct {
	CommandService command.CommandService
}

func (h *WhoamiCommand) Execute(_ []string) error {
	token, err := core.LoadToken()

	if err != nil {
		return err
	}

	apiLink := h.CommandService.ApiLink
	repo := moodle.NewMoodleRepository(apiLink)

	userInfo, err := repo.GetUserInfo(token)
	if err != nil {
		log.Fatalf("Error getting user info: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	row := fmt.Sprintf("User Info:\nBarcode \t %s\nFull Name \t %s\nUser ID \t %s\n", userInfo.Barcode, userInfo.FullName, userInfo.UserID)

	if _, err := fmt.Fprintln(tw, row); err != nil {
		return fmt.Errorf("Error writing to tabwriter: %v", err)
	}
	if err := tw.Flush(); err != nil {
		return fmt.Errorf("Error flushing tabwriter: %v", err)
	}

	return nil
}

func (h *WhoamiCommand) Description() string {
	return "Show current user info"
}
