package command

type CommandService struct {
	CommandRepository *CommandRepository
	ApiLink           string
}

func NewCommandService(repository *CommandRepository, apiLink string) CommandService {
	commandRepository := repository
	return CommandService{CommandRepository: commandRepository, ApiLink: apiLink}
}

func (cs *CommandService) RegisterCommand(name string, cmd Command) {
	cs.CommandRepository.RegisterCommand(name, cmd)
}

func (cs *CommandService) GetAvailableCommands() map[string]Command {
	return cs.CommandRepository.GetCommands()
}

func (cs *CommandService) GetCommand(name string) (Command, bool) {
	return cs.CommandRepository.GetCommand(name)
}
