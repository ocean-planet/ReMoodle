package command

type CommandRepository struct {
	commands map[string]Command
}

func NewCommandRepository() *CommandRepository {
	return &CommandRepository{
		commands: make(map[string]Command),
	}
}

func (cr *CommandRepository) RegisterCommand(name string, cmd Command) {
	cr.commands[name] = cmd
}

func (cr *CommandRepository) GetCommands() map[string]Command {
	return cr.commands
}

func (cr *CommandRepository) GetCommand(name string) (Command, bool) {
	cmd, found := cr.commands[name]
	return cmd, found
}
