package main

import (
	"github.com/ocean-planet/ReMoodle/internal/app/moodle"
)

func main() {
	//ctx := context.Background()
	// Инициализация репозитория и сервиса
	repo := moodle.NewMoodleRepository("https://moodle.astanait.edu.kz/webservice/rest/server.php?wstoken=")
	service := moodle.NewMoodleService(repo)

	service.GetDeadlines("")

	// Настройка маршрутов HTTP
	//http.HandleFunc("/user", moodle.GetUserHandler)
	//http.HandleFunc("/courses", moodle.GetUserCoursesHandler)
	//http.HandleFunc("/deadlines", moodle.GetDeadlinesHandler)

	// Запуск сервера
	//port := 8080
	//fmt.Printf("Server is running on port %d\n", port)
	//http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
