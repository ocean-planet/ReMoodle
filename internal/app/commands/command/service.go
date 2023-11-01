package command

type Service struct {
	CommandRepository *Repository
	ApiLink           string
}

func NewCommandService(repository *Repository, apiLink string) Service {
	commandRepository := repository
	return Service{CommandRepository: commandRepository, ApiLink: apiLink}
}

func (cs *Service) RegisterCommand(name string, cmd Command) {
	cs.CommandRepository.RegisterCommand(name, cmd)
}

func (cs *Service) GetAvailableCommands() map[string]Command {
	return cs.CommandRepository.GetCommands()
}

func (cs *Service) GetCommand(name string) (Command, bool) {
	return cs.CommandRepository.GetCommand(name)
}
