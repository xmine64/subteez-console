package config

import (
	"fmt"
	"subteez/config"
	"subteez/messages"
)

func dump(args []string, config config.Config) error {
	if config.IsScriptMode() {
		fmt.Printf(
			messages.DumpConfig,
			config.GetServer(),
			config.IsInteractive(),
			config.IsScriptMode(),
			config.GetLanguageFilters(),
		)
	} else {
		fmt.Printf(
			messages.DumpConfigHumanReadable,
			config.GetServer(),
			config.IsInteractive(),
			config.IsScriptMode(),
		)

		for _, language := range config.GetLanguageFilters() {
			fmt.Printf(" %s\n", language.GetTitle())
		}
		fmt.Println()
	}

	return nil
}
