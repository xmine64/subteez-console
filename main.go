package main

import (
	"fmt"
	"log"
	"os"
	"subteez/commands"
	"subteez/config"
	"subteez/subteez"
)

func main() {
	appConfig := config.NewConfigFile("subteez.config")
	if appConfig.Load() != nil {
		appConfig.SetServer("https://subteez1.herokuapp.com")
		appConfig.SetLanguageFilters(subteez.Languages)
		if appConfig.Save() != nil {
			fmt.Println("Note: Error in saving default config")
		}
	}

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	commands.AllCommands["help"] = commands.HelpCommand

	command, exists := commands.AllCommands[os.Args[1]]
	if !exists {
		log.Fatalf("command %s not found", os.Args[1])
	}
	err := command.Main(os.Args[1:], appConfig)
	if err != nil {
		log.Fatal(err)
	}

}
