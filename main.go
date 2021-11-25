package main

import (
	"flag"
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
		err := appConfig.Save()
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(os.Args) < 2 {
		log.Fatal(messages.NotEnoughArguments)
	}

	interactiveFlag := flag.Bool("interactive", false, "Enable interactive mode")
	helpFlag := flag.Bool("help", false, "Get help")

	flag.Parse()

	if *interactiveFlag {
		appConfig.SetInteractive(true)
	}

	if *helpFlag {
		if err := commands.HelpCommand.Main(flag.Args(), appConfig); err != nil {
			log.Fatal(err)
		}
		return
	}

	command, exists := commands.AllCommands[flag.Args()[0]]
	if !exists {
		log.Fatalf(messages.CommandNotFound, flag.Args()[0])
	}
	if err := command.Main(flag.Args(), appConfig); err != nil {
		log.Fatal(err)
	}

}
