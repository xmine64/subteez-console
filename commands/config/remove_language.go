package config

import (
	"errors"
	"subteez/config"
	serr "subteez/errors"
	"subteez/messages"
	"subteez/subteez"
)

func removeLanguage(args []string, config config.Config) error {
	if len(args) < 2 {
		return errors.New(messages.NotEnoughArguments)
	}

	if language, err := subteez.ParseLanguage(args[1]); err == nil {
		if err := config.RemoveLanguageFilter(language); err == nil {
			return &serr.ConfigChanged{}
		} else {
			return err
		}
	} else {
		return err
	}
}
