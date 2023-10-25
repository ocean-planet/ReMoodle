package main

import (
	"fmt"
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
	"log"
)

func main() {
	apiToken := ""

	apiLink := "https://moodle.astanait.edu.kz/webservice/rest/server.php?wstoken=" + apiToken
	repo := moodle.NewMoodleRepository(apiLink)

	// Replace "yourtoken" with your actual Moodle API token.

	// Example 1: Get user information
	userInfo, err := repo.GetUserInfo(apiToken)
	if err != nil {
		log.Fatalf("Error getting user info: %v", err)
	}
	fmt.Printf("User Info:\nBarcode: %s\nFull Name: %s\nUser ID: %s\n", userInfo.Barcode, userInfo.FullName, userInfo.UserID)

	// Example 2: Get user's courses
	courses, err := repo.GetUserAllCourses(apiToken)
	if err != nil {
		log.Fatalf("Error getting user courses: %v", err)
	}
	fmt.Println("\nUser Courses:")
	for _, course := range courses {
		fmt.Printf("Course Name: %s\nEnrolled User Count: %d\n", course.Name, course.EnrolledUserCount)
	}

	// Example 3: Get upcoming deadlines
	deadlines, err := repo.GetDeadlines(apiToken)
	if err != nil {
		log.Fatalf("Error getting deadlines: %v", err)
	}
	fmt.Println("\nUpcoming Deadlines:")
	for _, deadline := range deadlines {
		fmt.Printf("Course Name: %s\nDeadline Name: %s\nRemaining Time: %d\n", deadline.CourseName, deadline.DeadlineName, deadline.Remaining)
	}

	//TODO Convert GPT ideas to real code
	//// Define the Moodle API URL and token
	//moodleAPILink := "https://moodle.astanait.edu.kz/webservice/rest/server.php"
	//token := "your-moodle-token"
	//
	//// Create a Moodle repository
	//repo := moodle.NewMoodleRepository(moodleAPILink)
	//
	//// Define the local file path, file name, and the URL for file upload
	//filePath := "/path/to/your/file.pdf" // Replace with the actual file path
	//fileName := "file.pdf"
	//uploadURL := "https://your-moodle-site/upload/file.php" // Replace with the actual upload URL
	//
	//// Upload the file
	//response, err := repo.UploadFile(token, filePath, fileName, uploadURL)
	//if err != nil {
	//	log.Fatalf("File upload failed: %v", err)
	//}
	//
	//// Check the response
	//if exception, ok := response["exception"].(string); ok {
	//	fmt.Printf("File upload failed with exception: %s\n", exception)
	//} else {
	//	fmt.Println("File uploaded successfully.")
	//}

}
