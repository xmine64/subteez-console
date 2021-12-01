// main of "set-filter" subcommand

package config_command

import (
	"subteez/config"
	"subteez/errors"
	"subteez/subteez"
)

func setFilter(args []string, config config.Config) error {
	if len(args) < 2 {
		return errors.ErrNotEnoughArguments
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

	return errors.ErrConfigChanged
}
