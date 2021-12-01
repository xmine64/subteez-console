package config

import (
	"errors"
	"fmt"
	"strconv"
	"subteez/config"
	serr "subteez/errors"
	"subteez/messages"
)

func set(args []string, config config.Config) error {
	if len(args) < 3 {
		return errors.New(messages.NotEnoughArguments)
	}

	switch args[1] {
	case "server":
		config.SetServer(args[2])
		return &serr.ConfigChanged{}
	case "interactive":
		if boolValue, err := strconv.ParseBool(args[2]); err == nil {
			config.SetInteractive(boolValue)
			return &serr.ConfigChanged{}
		} else {
			return err
		}
	default:
		return fmt.Errorf(messages.OptionNotFound, args[1])
	}
}
