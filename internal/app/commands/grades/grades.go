package grades

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

type GradesCommand struct {
	CommandService command.CommandService
}

func (g *GradesCommand) Description() string {
	return "shows grades for a course (takes courseID as an argument)"
}


func (g *GradesCommand) Execute(args []string) error {
	token, tokenErr := core.LoadToken()

	if (tokenErr != nil) {
		return tokenErr
	}

	if len(args) < 1 {
		fmt.Println("Usage: courses [courseID]")
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

		if (len(grade.GradeName) < 2) {
			continue
		}

		row := fmt.Sprintf("📋  %s \t  %.2f",
            grade.GradeName,
			grade.Value,
        )

        fmt.Fprintln(tw, row)
	}
    tw.Flush()
	fmt.Println()

	return nil
}
