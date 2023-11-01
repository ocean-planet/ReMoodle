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

	apiLink := "https://moodle.astanait.edu.kz/webservice/rest/server.php"
	repo := moodle.NewMoodleRepository(apiLink)

	userInfo, err := repo.GetUserInfo(token)
	if err != nil {
		log.Fatalf("Error getting user info: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	row := fmt.Sprintf("User Info:\nBarcode \t %s\nFull Name \t %s\nUser ID \t %s\n", userInfo.Barcode, userInfo.FullName, userInfo.UserID)

	fmt.Fprintln(tw, row)
	tw.Flush()

	return nil
}

func (h *WhoamiCommand) Description() string {
	return "Show current user info"
}
