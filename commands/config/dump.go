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
		config.GetLanguageFilters(),
		config.IsInteractive(),
	)
	return nil
}
