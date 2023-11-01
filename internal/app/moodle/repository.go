package moodle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type MoodleRepository struct {
	MoodleAPILink string
}

func NewMoodleRepository(apiLink string) *MoodleRepository {
	return &MoodleRepository{
		MoodleAPILink: apiLink,
	}
}

func (r *MoodleRepository) GetUserInfo(token string) (*User, error) {
	url := fmt.Sprintf("%s?wstoken=%s&wsfunction=core_webservice_get_site_info&moodlewsrestformat=json", r.MoodleAPILink, token)
	fmt.Println("URL:", url) // Debugging: Print the URL

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}(resp.Body)

	fmt.Println("Response Status Code:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("moodle API request failed with status code %d", resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}

	user := &User{
		Barcode:  response["username"].(string),
		FullName: response["fullname"].(string),
		UserID:   strconv.FormatFloat(response["userid"].(float64), 'f', -1, 64),
	}

	return user, nil
}

func (r *MoodleRepository) GetUserAllCourses(token string) ([]Course, error) {
	user, err := r.GetUserInfo(token)
	if err != nil {
		return nil, err
	}
	userID := user.UserID

	url := fmt.Sprintf("%s?wstoken=%s&wsfunction=core_enrol_get_users_courses&moodlewsrestformat=json&userid=%s", r.MoodleAPILink, token, userID)
	fmt.Println("URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}(resp.Body)

	fmt.Println("Response Status Code:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("moodle API request failed with status code %d", resp.StatusCode)
	}

	var response []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}

	courses := make([]Course, len(response))

	for i, courseData := range response {

		course := Course{
			ID:                int(courseData["id"].(float64)),
			Name:              courseData["displayname"].(string),
			EnrolledUserCount: int(courseData["enrolledusercount"].(float64)),
			Category:          strconv.FormatFloat(courseData["category"].(float64), 'f', -1, 64),
			EndDate: 		   int64(courseData["enddate"].(float64)),
			// You may need to handle other fields like 'completed', 'start_date', and 'end_date' based on your response structure.
		}
		courses[i] = course
	}

	return courses, nil
}

func (r *MoodleRepository) GetDeadlines(token string) ([]Deadline, error) {
	url := fmt.Sprintf("%s?wstoken=%s&wsfunction=core_calendar_get_calendar_upcoming_view&moodlewsrestformat=json", r.MoodleAPILink, token)
	fmt.Println("URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}(resp.Body)

	fmt.Println("Response Status Code:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("moodle API request failed with status code %d", resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}

	events, ok := response["events"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for deadlines")
	}

	deadlines := make([]Deadline, len(events))

	for i, event := range events {
		eventData, ok := event.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid event data format")
		}

		deadline := Deadline{
			ID:           int(eventData["id"].(float64)),
			CourseName:   eventData["course"].(map[string]interface{})["shortname"].(string),
			DeadlineName: eventData["name"].(string),
			Remaining:    int64(eventData["timestart"].(float64)),
		}
		deadlines[i] = deadline
	}

	return deadlines, nil
}

// TODO Debug this code
func (r *MoodleRepository) UploadFile(token, filePath, fileName, url string) (map[string]interface{}, error) {

	client := &http.Client{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}(resp.Body)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
