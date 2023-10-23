package moodle

type MoodleService struct {
	Repository *MoodleRepository
}

func NewMoodleService(repo *MoodleRepository) *MoodleService {
	return &MoodleService{
		Repository: repo,
	}
}

func (s *MoodleService) GetUserInfo(token string) (*User, error) {
	return s.Repository.GetUserInfo(token)
}

func (s *MoodleService) GetUserAllCourses(token string) ([]Course, error) {
	return s.Repository.GetUserAllCourses(token)
}

func (s *MoodleService) GetDeadlines(token string) ([]Deadline, error) {
	return s.Repository.GetDeadlines(token)
}

//// HTTP Handlers
//func GetUserHandler(w http.ResponseWriter, r *http.Request) {
//	// Обработчик для получения информации о пользователе
//}
//
//func GetUserCoursesHandler(w http.ResponseWriter, r *http.Request) {
//	// Обработчик для получения всех курсов пользователя
//}
//
//func GetDeadlinesHandler(w http.ResponseWriter, r *http.Request) {
//	// Обработчик для получения сроков
//}
