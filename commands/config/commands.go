package config

import "subteez/config"

var commands = map[string]func(args []string, config config.Config) error{
	"show": dump,
	"set":  set,
}
