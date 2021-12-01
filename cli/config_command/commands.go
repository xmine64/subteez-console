// subcommands definitions for config command

package config_command

import "subteez/config"

var commands = map[string]func(args []string, config config.Config) error{
	"show":       dump,
	"set":        set,
	"add-filter": addFilter,
	"rm-filter":  removeFilter,
	"set-filter": setFilter,
}
