package config

import (
	"subteez/config"
	"subteez/errors"
	"subteez/interactive"
)

func Main(args []string, config config.Config) error {
	if len(args) < 2 {
		if config.IsInteractive() {
			context := interactive.Context{}
			context.Initialize(config)
			context.NavigateToConfig()
			return context.Run()
		}
		return dump(nil, config)
	}

	command, exists := commands[args[1]]
	if !exists {
		return errors.ErrCommandNotFound(args[1])
	}
	return command(args[1:], config)
}
