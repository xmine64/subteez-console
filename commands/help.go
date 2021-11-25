package commands

import (
	"fmt"
	"subteez/config"
	"subteez/constants"
	"subteez/messages"
)

var HelpCommand = Command{
	Name:        "help",
	Description: "Get help about using this app",
	Main: func(args []string, _ config.Config) error {
		fmt.Printf(
			messages.HelpMessage,
			constants.Name,
			constants.VersionMajor,
			constants.VersionMinor,
			constants.VersionBuild,
			constants.ExeName,
		)

		for _, command := range AllCommands {
			fmt.Printf(messages.CommandRow, command.Name, command.Description)
		}
		return nil
	},
}
