package main

import (
	"os"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/help"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
)

func main() {
	commandRepository := command.NewCommandRepository()
	commandService := command.NewCommandService(commandRepository)
	myApp := core.NewApp(commandService)

	commandRepository.RegisterCommand("help", &help.HelpCommand{CommandService: commandService})

	if err := myApp.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
