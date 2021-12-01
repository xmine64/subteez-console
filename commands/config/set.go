package config

import (
	"errors"
	"fmt"
	"strconv"
	"subteez/config"
	"subteez/messages"
)

func set(args []string, config config.Config) error {
	if len(args) < 3 {
		return errors.New(messages.NotEnoughArguments)
	}

	switch args[1] {
	case "server":
		config.SetServer(args[2])
		return nil
	case "interactive":
		if boolValue, err := strconv.ParseBool(args[2]); err == nil {
			config.SetInteractive(boolValue)
			return nil
		} else {
			return err
		}
	default:
		return fmt.Errorf(messages.OptionNotFound, args[1])
	}
}
