package deadlines

import (
	"fmt"
	"strings"
	"time"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

type DeadlineCommand struct {
	CommandService command.CommandService
}

func (d *DeadlineCommand) Execute(args []string) error {
	token, tokenErr := core.LoadToken()

	if (tokenErr != nil) {
		return tokenErr
	}

	moodleRepository := moodle.NewMoodleRepository("https://moodle.astanait.edu.kz/webservice/rest/server.php")
	moodleService := moodle.NewMoodleService(moodleRepository)

	deadlines, err := moodleService.GetDeadlines(token)

	if err != nil {
		fmt.Println("Moodle is currently unavailable, try again later")
		return err
	}

	if len(deadlines) < 1 {
		fmt.Println("Congratulations! There is no deadlines at the moment for you")
		return nil
	}

	fmt.Println("> ------- Current Deadlines ------- <")

	for _, deadline := range deadlines {

		if strings.Contains(strings.ToLower(deadline.DeadlineName), "attendance") {
			continue
		}

		fmt.Println("> " + deadline.DeadlineName + " | Date: " + GetDateString(deadline.Remaining) + " | Time left: " + getRemainingString(deadline.Remaining))
	}

	return nil
}

func (d *DeadlineCommand) Description() string {
	return "shows all active deadlines"
}

func GetDateString(unixtimestamp int64) string {
	deadlineTime := time.Unix(unixtimestamp, 0)
	return deadlineTime.Format("2006-01-02 15:04:05")
}

func getRemainingString(unixtimestamp int64) string {	
	finalString := "Time left: "

	currentTime := time.Now()
	deadlineTime := time.Unix(unixtimestamp, 0)
	timeDelta := deadlineTime.Sub(currentTime)

	if (timeDelta.Hours() > 24) {
		finalString += string(int32(timeDelta.Hours()) / 24) + " days, "
	}

	finalString += string(int32(timeDelta.Hours()) % 24) + "h "
	finalString += string(int32(timeDelta.Minutes())) + "m "
	finalString += string(int32(timeDelta.Seconds())) + "s "

	return finalString
}
