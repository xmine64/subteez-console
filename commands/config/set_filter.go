package config

import (
	"errors"
	"subteez/config"
	serr "subteez/errors"
	"subteez/messages"
	"subteez/subteez"
)

func setFilter(args []string, config config.Config) error {
	if len(args) < 2 {
		return errors.New(messages.NotEnoughArguments)
	}

	config.ClearLanguageFilters()

	for _, arg := range args[1:] {
		if language, err := subteez.ParseLanguage(arg); err == nil {
			if err := config.AddLanguageFilter(language); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return &serr.ConfigChanged{}
}
