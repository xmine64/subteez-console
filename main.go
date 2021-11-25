package main

import (
	"fmt"
	"log"
	"os"
	"subteez/commands"
	"subteez/config"
	"subteez/constants"
	"subteez/messages"
	"subteez/subteez"
)

func main() {
	appConfig := config.NewConfigFile(constants.ConfigFileName)
	if appConfig.Load() != nil {
		appConfig.SetServer(constants.DefaultServer)
		appConfig.SetLanguageFilters(subteez.Languages)
		if appConfig.Save() != nil {
			fmt.Println("Note: Error in saving default config")
		}
	}

	if len(os.Args) < 2 {
		log.Fatal(messages.NotEnoughArguments)
	}

	commands.AllCommands["help"] = commands.HelpCommand

	command, exists := commands.AllCommands[os.Args[1]]
	if !exists {
		log.Fatalf(messages.CommandNotFound, os.Args[1])
	}
	err := command.Main(os.Args[1:], appConfig)
	if err != nil {
		log.Fatal(err)
	}

}
