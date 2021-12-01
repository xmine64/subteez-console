// main of config command

package config_command

import (
	"subteez/config"
	"subteez/errors"
	"subteez/tui"
)

func Main(args []string, config config.Config) error {
	// if no command given, run interactive mode if it's enabled, or dump configurations
	if len(args) < 2 {
		if config.IsInteractive() {
			context := tui.Context{}
			context.Initialize(config)
			context.NavigateToConfig()
			return context.Run()
		}
		return dump(nil, config)
	}

	// run given command and return its error
	if command, exists := commands[args[1]]; exists {
		return command(args[1:], config)
	} else {
		return errors.ErrCommandNotFound(args[1])
	}
}
