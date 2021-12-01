package errors

import (
	"errors"
	"fmt"
	"subteez/messages"
)

var ErrNotEnoughArguments = errors.New(messages.NotEnoughArguments)

func ErrCommandNotFound(value string) error {
	return fmt.Errorf(messages.CommandNotFound, value)
}

var ErrInteractiveAndScript = errors.New(messages.InteractiveAndScript)
