package command

type CommandService struct {
	CommandRepository *CommandRepository
}

func NewCommandService(repository *CommandRepository) CommandService {
	commandRepository := repository
	return CommandService{CommandRepository: commandRepository}
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
