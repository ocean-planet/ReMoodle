package moodle

type MoodleRepository struct {
	MoodleAPILink string
}

func NewMoodleRepository(apiLink string) *MoodleRepository {
	return &MoodleRepository{
		MoodleAPILink: apiLink,
	}
}

func (r *MoodleRepository) GetUserInfo(token string) (*User, error) {
	// Реализация получения информации о пользователе из Moodle API
	return nil, nil
}

func (r *MoodleRepository) GetUserAllCourses(token string) ([]Course, error) {
	// Реализация получения всех курсов пользователя из Moodle API
	return nil, nil
}

func (r *MoodleRepository) GetDeadlines(token string) ([]Deadline, error) {
	// Реализация получения сроков из Moodle API
	return nil, nil
}
