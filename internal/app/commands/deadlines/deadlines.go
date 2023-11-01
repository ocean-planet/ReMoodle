package deadlines

import (
	"fmt"
	"strconv"
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

	fmt.Println("Current Deadlines:")

	for _, deadline := range deadlines {

		if strings.Contains(strings.ToLower(deadline.DeadlineName), "attendance") {
			continue
		}

		fmt.Println("> 📋 " + deadline.DeadlineName + " | 📚 " + strings.Split(deadline.CourseName, " | ")[0] + " | 📅 Date: " + GetDateString(deadline.Remaining) + " | ⌚ Time left: " + GetRemainingString(deadline.Remaining))
	}
	fmt.Println()

	return nil
}

func (d *DeadlineCommand) Description() string {
	return "shows all active deadlines"
}

func GetDateString(unixtimestamp int64) string {
	deadlineTime := time.Unix(unixtimestamp, 0)
	return deadlineTime.Format("2006-01-02 15:04:05")
}

func GetRemainingString(unixtimestamp int64) string {	
	finalString := ""

	currentTime := time.Now()
	deadlineTime := time.Unix(unixtimestamp, 0)
	timeDelta := deadlineTime.Sub(currentTime)

	if timeDelta.Hours() > 24 {
		finalString += strconv.Itoa(int(timeDelta.Hours())/24) + " days, "
	}

	finalString += strconv.Itoa(int(timeDelta.Hours())%24) + "h "
	finalString += strconv.Itoa(int(timeDelta.Minutes())%60) + "m "
	finalString += strconv.Itoa(int(timeDelta.Seconds())%60) + "s "

	return finalString
}
