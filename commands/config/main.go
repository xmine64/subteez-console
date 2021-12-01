package config

import (
	"subteez/config"
	"subteez/errors"
)

func Main(args []string, config config.Config) error {
	if len(args) < 2 {
		return dump(nil, config)
	}

	command, exists := commands[args[1]]
	if !exists {
		return errors.ErrCommandNotFound(args[1])
	}
	return command(args[1:], config)
}
