package command

import "errors"

var (
	ErrNotEnoughArguments = errors.New("Not enough arguments for this command")
)
