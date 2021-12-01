package config

import (
	"errors"
	"subteez/config"
	serr "subteez/errors"
	"subteez/messages"
	"subteez/subteez"
)

func addFilter(args []string, config config.Config) error {
	if len(args) < 2 {
		return errors.New(messages.NotEnoughArguments)
	}

	if language, err := subteez.ParseLanguage(args[1]); err == nil {
		if err := config.AddLanguageFilter(language); err == nil {
			return &serr.ConfigChanged{}
		} else {
			return err
		}
	} else {
		return err
	}
}
