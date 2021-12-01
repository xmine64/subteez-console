// main of "show" subcommand

package config_command

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
			true,
			config.GetLanguageFilters(),
		)
	} else {
		fmt.Printf(
			messages.DumpConfigHumanReadable,
			config.GetServer(),
			config.IsInteractive(),
			false,
		)

		for _, language := range config.GetLanguageFilters() {
			fmt.Printf(" %s\n", language.GetTitle())
		}
		fmt.Println()
	}

	return nil
}
