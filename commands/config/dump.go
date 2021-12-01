package config

import (
	"fmt"
	"subteez/config"
	"subteez/messages"
)

func dump(args []string, config config.Config) error {
	fmt.Printf(
		messages.DumpConfig,
		config.GetServer(),
		config.IsInteractive(),
	)

	for _, language := range config.GetLanguageFilters() {
		fmt.Printf(" %s\n", language.GetTitle())
	}
	fmt.Println()

	return nil
}
