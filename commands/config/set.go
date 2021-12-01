package config

import (
	"strconv"
	"subteez/config"
	"subteez/errors"
)

func set(args []string, config config.Config) error {
	if len(args) < 3 {
		return errors.ErrNotEnoughArguments
	}

	switch args[1] {
	case "server":
		config.SetServer(args[2])
		return errors.ErrConfigChanged
	case "interactive":
		if boolValue, err := strconv.ParseBool(args[2]); err == nil {
			config.SetInteractive(boolValue)
			return errors.ErrConfigChanged
		} else {
			return err
		}
	default:
		return errors.ErrConfigOptionNotFound(args[1])
	}
}
