package errors

import (
	"errors"
	"subteez/messages"
)

var ErrEmptyID = errors.New(messages.EmptyID)

var ErrNoFileFound = errors.New(messages.NoFileFound)
