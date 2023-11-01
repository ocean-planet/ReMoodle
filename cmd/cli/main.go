package main

import (
	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/courses"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/deadlines"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/grades"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/help"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/login"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/logout"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/whoami"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
	"os"
)

func main() {
	commandRepository := command.NewCommandRepository()
	commandService := command.NewCommandService(commandRepository, "https://moodle.astanait.edu.kz/webservice/rest/server.php")
	myApp := core.NewApp(commandService)

	commandRepository.RegisterCommand("help", &help.Command{CommandService: commandService})
	commandRepository.RegisterCommand("login", &login.Command{CommandService: commandService})
	commandRepository.RegisterCommand("logout", &logout.Command{CommandService: commandService})
	commandRepository.RegisterCommand("whoami", &whoami.Command{CommandService: commandService})
	commandRepository.RegisterCommand("deadlines", &deadlines.Command{CommandService: commandService})
	commandRepository.RegisterCommand("courses", &courses.Command{CommandService: commandService})
	commandRepository.RegisterCommand("grades", &grades.Command{CommandService: commandService})

	if err := myApp.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
