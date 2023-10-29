package deadlines

import (
	"fmt"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

type DeadlineCommand struct {
	CommandService command.CommandService
}

func (d *DeadlineCommand) Execute(args []string) error {
	if len(args) > 1 {
		fmt.Println("Usage: remoodle -d 'moodletoken'")
	}

	apiToken := args[0]

	moodleRepository := moodle.NewMoodleRepository(apiToken)
	moodleService := moodle.NewMoodleService(moodleRepository)

	userInfo, tokenError := moodleRepository.GetUserInfo(apiToken)
	if tokenError != nil {
		fmt.Println("Token is invalid!")
		return nil
	}

	deadlines, err := moodleService.GetDeadlines(apiToken)

	if err != nil {
		fmt.Println("Moodle API is currently unavailable, try again later")
		return err
	}

	if len(deadlines) < 1 {
		fmt.Println("Congratulations! There is no available deadlines")
		return nil
	}

	fmt.Println("> ------- < Current Deadlines > ------- <")

	for _, deadline := range deadlines {
		fmt.Println("> " + deadline.DeadlineName + " | " + deadline.CourseName)
		// fmt.Printf("  Date: " + deadline.Remaining + " | Time left: " + deadline.Remaining)
	}

	return nil
}

func (d *DeadlineCommand) Description() string {
	return "shows all active deadlines"
}