package config

import (
	"fmt"
	"subteez/config"
	"subteez/messages"
)

func Main(args []string, cfg config.Config) error {
	fmt.Printf(
		messages.DumpConfig,
		cfg.GetServer(),
		cfg.GetLanguageFilters(),
		cfg.IsInteractive(),
	)
	return nil
}
