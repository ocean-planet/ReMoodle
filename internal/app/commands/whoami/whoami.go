package whoami

import (
	"fmt"
	"log"

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
	fmt.Printf("User Info:\n> Barcode: %s\n> Full Name: %s\n> User ID: %s\n", userInfo.Barcode, userInfo.FullName, userInfo.UserID)

	return nil
}

func (h *WhoamiCommand) Description() string {
	return "Show current user info"
}
