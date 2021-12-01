package main

import (
	"flag"
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
		err := appConfig.Save()
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(os.Args) < 2 {
		showHelp()
		log.Fatal(messages.NotEnoughArguments)
	}

	interactiveFlag := flag.Bool("interactive", false, "")
	helpFlag := flag.Bool("help", false, "")

	flag.Parse()

	if *interactiveFlag {
		appConfig.SetInteractive(true)
	}

	if *helpFlag {
		showHelpTopic(flag.Arg(0))
		return
	}

	command, exists := commands.AllCommands[flag.Args()[0]]
	if !exists {
		showHelp()
		log.Fatalf(messages.CommandNotFound, flag.Args()[0])
	}
	if err := command.Main(flag.Args(), appConfig); err != nil {
		log.Fatal(err)
	}

}

func showHelp() {
	fmt.Printf(
		messages.HelpMessage,
		constants.Name,
		constants.VersionMajor,
		constants.VersionMinor,
		constants.VersionBuild,
		constants.ExeName,
	)

	for _, command := range commands.AllCommands {
		fmt.Printf(messages.CommandRow, command.Name, command.Description)
	}
}

func showHelpTopic(topic string) {
	command, exists := commands.AllCommands[topic]

	if !exists {
		showHelp()
		log.Fatalf(messages.TopicNotFound, topic)
	}

	fmt.Println(command.HelpTopic)
}
