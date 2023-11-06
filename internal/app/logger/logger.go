package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger interface {
	Log(message string)
}

type DefaultLogger struct {
	log *log.Logger
}

func (l *DefaultLogger) Log(message string) {
	l.log.Println(message)
}

type Decorator interface {
	Log(message string)
}

type BaseDecorator struct {
	Logger Logger
}

func (d *BaseDecorator) Log(message string) {
	d.Logger.Log(message)
}

type LogDecorator struct {
	*BaseDecorator
	Prefix string
}

func (d *LogDecorator) Log(message string) {
	fmt.Printf("%s: %s\n", d.Prefix, message)
}

var instance *DefaultLogger
var once sync.Once

func GetInstance() *DefaultLogger {
	once.Do(func() {
		file, err := os.Create("remoodle.log")
		if err != nil {
			log.Fatalf("Error creating log file: %v", err)
		}

		instance = &DefaultLogger{log: log.New(file, "", log.Ldate|log.Ltime)}
	})
	return instance
}

func LogWithPrefix(prefix string, message ...string) {
	GetInstance().log.Println(prefix, ":", message)
}

func LogWithDecorator(decorator Decorator, message string) {
	decorator.Log(message)
}
