package deadlines

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

type DeadlineCommand struct {
	CommandService command.CommandService
}

func (d *DeadlineCommand) Execute(_ []string) error {
	token, tokenErr := core.LoadToken()

	if tokenErr != nil {
		return tokenErr
	}

	moodleRepository := moodle.NewMoodleRepository(d.CommandService.ApiLink)
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

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	for _, deadline := range deadlines {

		if strings.Contains(strings.ToLower(deadline.DeadlineName), "attendance") {
			continue
		}

		row := fmt.Sprintf("📋  %s\t 📚  %s\t 📅 Date: %s\t ⌚ Time left: %s",
			deadline.DeadlineName,
			strings.Split(deadline.CourseName, " | ")[0],
			GetDateString(deadline.Remaining),
			GetRemainingString(deadline.Remaining),
		)

		if _, err := fmt.Fprintln(tw, row); err != nil {
			return fmt.Errorf("Error writing to tabwriter: %v", err)
		}
	}
	if err := tw.Flush(); err != nil {
		return fmt.Errorf("Error flushing tabwriter: %v", err)
	}

	fmt.Println()

	return nil
}

func (d *DeadlineCommand) Description() string {
	return "Shows all active deadlines"
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
