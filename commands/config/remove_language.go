package config

import (
	"errors"
	"subteez/config"
	"subteez/messages"
	"subteez/subteez"
)

func removeLanguage(args []string, config config.Config) error {
	if len(args) < 2 {
		return errors.New(messages.NotEnoughArguments)
	}

	if language, err := subteez.ParseLanguage(args[1]); err == nil {
		return config.RemoveLanguageFilter(language)
	} else {
		return err
	}
}
