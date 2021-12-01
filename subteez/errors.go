package subteez

import (
	"errors"
	"fmt"
)

// error when requested remote resource (like file, or subtitle details) not found and server responded with code 404
var ErrNotFound = errors.New("requested resource not found")

// error when a bad request sent to server and server responded with code 400
var ErrBadRequest = errors.New("bad request")

// when an error occured on server and server responded with code 500
var ErrServer = errors.New("server error")

// when an error occured on client
type ClientError struct {
	message string
}

func (e *ClientError) Error() string {
	return "client error: " + e.message
}

func ErrClientError(message string) error {
	return &ClientError{message: message}
}

func ErrUnhandledResponse(status string) error {
	return ErrClientError(fmt.Sprintf(`server responded with "%s" and it's not handled`, status))
}
