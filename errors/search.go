package errors

import (
	"errors"
	"subteez/messages"
)

var ErrEmptyQuery = errors.New(messages.EmptyQuery)

var ErrNoSearchResult = errors.New(messages.NoSearchResult)
