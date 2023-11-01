package logger

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	log *log.Logger
}

var instance *Logger
var once sync.Once

func GetInstance() *Logger {
	once.Do(func() {
		instance = createLogger()
	})
	return instance
}

func createLogger() *Logger {
	file, err := os.Create("remoodle.log")
	if err != nil {
		log.Fatalf("Error creating log file: %v", err)
	}

	return &Logger{log: log.New(file, "", log.Ldate|log.Ltime)}
}

//func (l *Logger) Log(message string) {
//	l.log.Println(message)
//}

func Log(message ...string) {
	GetInstance().log.Println(message)
}
