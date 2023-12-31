package grades

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

type Command struct {
	CommandService command.Service
}

func (c *Command) Description() string {
	return "shows grades for a course (takes courseID as an argument)"
}

func (c *Command) Execute(args []string) error {
	token, tokenErr := core.LoadToken()

	if tokenErr != nil {
		return tokenErr
	}

	if len(args) < 1 {
		fmt.Println("Usage: courses [courseID]")
		return command.ErrNotEnoughArguments
	}

	courseID := args[0]

	moodleRepository := moodle.NewMoodleRepository("https://moodle.astanait.edu.kz/webservice/rest/server.php")
	moodleService := moodle.NewMoodleService(moodleRepository)

	grades, err := moodleService.GetUserCourseGrades(token, courseID)

	if err != nil {
		// fmt.Println("Moodle is currently unavailable, try again later")
		return err
	}

	if len(grades) < 1 {
		fmt.Println("You don't have grades in this course!")
		return nil
	}

	fmt.Printf("Grades for course with ID %s:\n", courseID)

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	for _, grade := range grades {

		if len(grade.GradeName) < 2 {
			continue
		}

		row := fmt.Sprintf("📋  %s \t  %.2f",
			grade.GradeName,
			grade.Value,
		)

		if _, err := fmt.Fprintln(tw, row); err != nil {
			return fmt.Errorf("error writing to tabwriter: %v", err)
		}
	}
	if err := tw.Flush(); err != nil {
		return fmt.Errorf("error flushing tabwriter: %v", err)
	}
	fmt.Println()

	return nil
}
