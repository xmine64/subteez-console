package errors

import (
	"fmt"
	"subteez/messages"
)

func ErrHelpTopicNotFound(value string) error {
	return fmt.Errorf(messages.TopicNotFound, value)
}
