package config

import (
	"fmt"
	"subteez/config"
	"subteez/messages"
)

func Main(args []string, config config.Config) error {
	if len(args) < 2 {
		return dump(nil, config)
	}

	command, exists := commands[args[1]]
	if !exists {
		return fmt.Errorf(messages.CommandNotFound, args[1])
	}
	return command(args[1:], config)
}
