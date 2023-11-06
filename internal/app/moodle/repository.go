package moodle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type Repository struct {
	MoodleAPILink string
}

func NewMoodleRepository(apiLink string) *Repository {
	return &Repository{
		MoodleAPILink: apiLink,
	}
}

func (r *Repository) GetUserInfo(token string) (*User, error) {
	url := fmt.Sprintf("%s?wstoken=%s&wsfunction=core_webservice_get_site_info&moodlewsrestformat=json", r.MoodleAPILink, token)
	//fmt.Println("URL:", url) // Debugging: Print the URL

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

	//fmt.Println("Response Status Code:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("moodle API request failed with status code %d", resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}

	if _, exists := response["username"]; exists {
		user := &User{
			Barcode:  response["username"].(string),
			FullName: response["fullname"].(string),
			UserID:   strconv.FormatFloat(response["userid"].(float64), 'f', -1, 64),
		}

		return user, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func (r *Repository) GetUserCourseGrades(token string, courseID string) ([]Grade, error) {
	// gradereport_user_get_grade_items

	user, err := r.GetUserInfo(token)
	if err != nil {
		return nil, err
	}
	userID := user.UserID

	url := fmt.Sprintf("%s?wstoken=%s&wsfunction=gradereport_user_get_grade_items&moodlewsrestformat=json&userid=%s&courseid=%s", r.MoodleAPILink, token, userID, courseID)
	//fmt.Println("URL:", url)

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

	//fmt.Println("Response Status Code:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("moodle API request failed with status code %d", resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}

	data, ok := response["usergrades"].([]interface{})

	if !ok || len(data) == 0 {
		fmt.Println("User's grades not of the expected type")
		return nil, err
	}

	userGrades, ok := data[0].(map[string]interface{})
	if !ok {
		return nil, err
	}

	gradeItems, ok := userGrades["gradeitems"].([]interface{})
	if !ok {
		fmt.Println("Grade items not of the expected type")
		return nil, err
	}

	grades := make([]Grade, len(gradeItems))

	for i, gradeDataInterface := range gradeItems {
		gradeData, ok := gradeDataInterface.(map[string]interface{})
		if !ok {
			fmt.Println("Grade data not of the expected type")
			return nil, errors.New("grade data not of the expected type")
		}

		if gradeData["graderaw"] == nil {
			gradeData["graderaw"] = 0.0
			gradeData["gradedategraded"] = 0.0
		}

		grade := Grade{
			GradeName:   gradeData["itemname"].(string),
			Value:       gradeData["graderaw"].(float64),
			MaxValue:    gradeData["grademax"].(float64),
			UpdatedDate: int64(gradeData["gradedategraded"].(float64)),

			// ID:                int(courseData["id"].(float64)),
			// Name:              courseData["displayname"].(string),
			// EnrolledUserCount: int(courseData["enrolledusercount"].(float64)),
			// Category:          strconv.FormatFloat(courseData["category"].(float64), 'f', -1, 64),
			// EndDate: 		   int64(courseData["enddate"].(float64)),
		}
		grades[i] = grade
	}

	return grades, nil
}

func (r *Repository) GetUserAllCourses(token string) ([]Course, error) {
	user, err := r.GetUserInfo(token)
	if err != nil {
		return nil, err
	}
	userID := user.UserID

	url := fmt.Sprintf("%s?wstoken=%s&wsfunction=core_enrol_get_users_courses&moodlewsrestformat=json&userid=%s", r.MoodleAPILink, token, userID)
	//fmt.Println("URL:", url)

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

	//fmt.Println("Response Status Code:", resp.Status)

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
			EndDate:           int64(courseData["enddate"].(float64)),
			// You may need to handle other fields like 'completed', 'start_date', and 'end_date' based on your response structure.
		}
		courses[i] = course
	}

	return courses, nil
}

func (r *Repository) GetDeadlines(token string) ([]Deadline, error) {
	url := fmt.Sprintf("%s?wstoken=%s&wsfunction=core_calendar_get_calendar_upcoming_view&moodlewsrestformat=json", r.MoodleAPILink, token)
	//fmt.Println("URL:", url)

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

	//fmt.Println("Response Status Code:", resp.Status)

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

func (r *Repository) UploadFile(token, filePath, fileName, url string) (map[string]interface{}, error) {

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
