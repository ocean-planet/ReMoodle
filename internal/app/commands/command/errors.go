package command

import "errors"

var (
	ErrNotEnoughArguments = errors.New("not enough arguments for this command")
)
