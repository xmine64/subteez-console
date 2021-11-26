package commands

import (
	"fmt"
	"subteez/config"
	"subteez/messages"
)

func mainConfig(args []string, cfg config.Config) error {
	fmt.Printf(
		messages.DumpConfig,
		cfg.GetServer(),
		cfg.GetLanguageFilters(),
		cfg.IsInteractive(),
	)
	return nil
}
