package courses

import (
	"fmt"
	"time"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

type CoursesCommand struct {
	CommandService command.CommandService
}

func (*CoursesCommand) Description() string {
	return "Shows current courses"
}

func (c *CoursesCommand) Execute(_ []string) error {
	token, tokenErr := core.LoadToken()

	if tokenErr != nil {
		return tokenErr
	}

	moodleRepository := moodle.NewMoodleRepository(c.CommandService.ApiLink)
	moodleService := moodle.NewMoodleService(moodleRepository)

	courses, err := moodleService.GetUserAllCourses(token)

	if err != nil {
		fmt.Println("Moodle API is currently unavailable, try again later")
		return err
	}

	if len(courses) < 1 {
		fmt.Println("Courses were not found")
		return nil
	}

	fmt.Println("Current courses:")

	currentUnixTime := time.Now().Unix()
	// relativeCourses := make([]moodle.Course, 0)

	for _, course := range courses {
		if course.EndDate > currentUnixTime {
			// relativeCourses = append(relativeCourses, course)
			fmt.Println("> 🆔 " + fmt.Sprint(course.ID) + " | 📚 " + course.Name)
		}
	}

	return nil
}
