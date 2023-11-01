package command

type Repository struct {
	commands map[string]Command
}

func NewCommandRepository() *Repository {
	return &Repository{
		commands: make(map[string]Command),
	}
}

func (cr *Repository) RegisterCommand(name string, cmd Command) {
	cr.commands[name] = cmd
}

func (cr *Repository) GetCommands() map[string]Command {
	return cr.commands
}

func (cr *Repository) GetCommand(name string) (Command, bool) {
	cmd, found := cr.commands[name]
	return cmd, found
}
