package commands

import (
	"fmt"
	"subteez/config"
)

var HelpCommand = Command{
	Name:        "help",
	Description: "Get help about using this app",
	Main: func(args []string, _ config.Config) error {
		fmt.Printf(`%s v%d.%d.%d

Usage: %s <Command> <Arguments>

Commands:

`,
			config.Name,
			config.VersionMajor,
			config.VersionMinor,
			config.VersionBuild,
			config.ExeName,
		)

		for _, command := range AllCommands {
			fmt.Printf("    %s:\t%s\n\n", command.Name, command.Description)
		}
		return nil
	},
}
