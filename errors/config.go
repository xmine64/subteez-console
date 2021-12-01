package errors

import (
	"fmt"
	"subteez/messages"
	"subteez/subteez"
)

type ConfigChanged struct{}

func (ConfigChanged) Error() string {
	return ""
}

// return this error to save configurations on exit
var ErrConfigChanged = &ConfigChanged{}

func ErrDuplicateLanguage(value subteez.Language) error {
	return fmt.Errorf(messages.DuplicateLanguage, value)
}

func ErrLanguageNotFound(value subteez.Language) error {
	return fmt.Errorf(messages.LanguageNotFound, value)
}

func ErrConfigOptionNotFound(value string) error {
	return fmt.Errorf(messages.ConfigOptionNotFound, value)
}
