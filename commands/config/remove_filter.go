package config

import (
	"subteez/config"
	"subteez/errors"
	"subteez/subteez"
)

func removeFilter(args []string, config config.Config) error {
	if len(args) < 2 {
		return errors.ErrNotEnoughArguments
	}

	if language, err := subteez.ParseLanguage(args[1]); err == nil {
		if err := config.RemoveLanguageFilter(language); err == nil {
			return errors.ErrConfigChanged
		} else {
			return err
		}
	} else {
		return err
	}
}
