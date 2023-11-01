package moodle

type Service struct {
	Repository *Repository
}

func NewMoodleService(repo *Repository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) GetUserInfo(token string) (*User, error) {
	return s.Repository.GetUserInfo(token)
}

func (s *Service) GetUserAllCourses(token string) ([]Course, error) {
	return s.Repository.GetUserAllCourses(token)
}

func (s *Service) GetDeadlines(token string) ([]Deadline, error) {
	return s.Repository.GetDeadlines(token)
}

func (s *Service) GetUserCourseGrades(token string, courseID string) ([]Grade, error) {
	return s.Repository.GetUserCourseGrades(token, courseID)
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
