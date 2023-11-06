package moodle

type User struct {
	Barcode  string `json:"barcode"`
	FullName string `json:"full_name"`
	UserID   string `json:"user_id"`
}

type Course struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	EnrolledUserCount int    `json:"enrolled_user_count"`
	Category          string `json:"category"`
	Completed         bool   `json:"completed"`
	StartDate         int64  `json:"start_date"`
	EndDate           int64  `json:"end_date"`
}

type Deadline struct {
	ID           int    `json:"id"`
	CourseName   string `json:"course_name"`
	DeadlineName string `json:"deadline_name"`
	Remaining    int64  `json:"remaining"`
	MarkDone     bool   `json:"mark_done"`
}

type Grade struct {
	GradeName   string
	Value       float64
	MaxValue    float64
	UpdatedDate int64
}
