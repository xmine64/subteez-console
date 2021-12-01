package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"subteez/commands"
	"subteez/config"
	"subteez/constants"
	"subteez/errors"
	"subteez/messages"
	"subteez/subteez"
)

func main() {
	configFile := config.NewConfigFile(constants.ConfigFileName)
	if configFile.Load() != nil {
		configFile.SetServer(constants.DefaultServer)
		configFile.SetLanguageFilters(subteez.Languages)
		err := configFile.Save()
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(os.Args) < 2 {
		showHelp()
		log.Fatal(errors.ErrNotEnoughArguments)
	}

	interactiveFlag := flag.Bool("interactive", false, "")
	scriptFlag := flag.Bool("script", false, "")
	helpFlag := flag.Bool("help", false, "")

	flag.Parse()

	if *helpFlag {
		showHelpTopic(flag.Arg(0))
		return
	}

	if *interactiveFlag && *scriptFlag {
		log.Fatal(errors.ErrInteractiveAndScript)
	}

	if *interactiveFlag {
		configFile.SetInteractive(true)
		configFile.SetScriptMode(false)
	} else if *scriptFlag {
		configFile.SetInteractive(false)
		configFile.SetScriptMode(true)
	}

	if !configFile.IsScriptMode() {
		log.Printf(messages.Version, constants.Name, constants.VersionMajor, constants.VersionMinor, constants.VersionBuild)
	}

	command, exists := commands.AllCommands[flag.Args()[0]]
	if !exists {
		if !configFile.IsScriptMode() {
			showHelp()
		}
		log.Fatal(errors.ErrCommandNotFound(flag.Args()[0]))
	}
	if err := command.Main(flag.Args(), configFile); err != nil {
		if _, ok := err.(*errors.ConfigChanged); ok {
			configFile.Save()
			log.Print(messages.ConfigFileSaved)
			return
		}
		log.Fatal(err)
	}
	if !configFile.IsInteractive() {
		log.Println(messages.Done)
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
	if topic == "" {
		showHelp()
		return
	}

	command, exists := commands.AllCommands[topic]

	if !exists {
		showHelp()
		log.Fatal(errors.ErrHelpTopicNotFound(topic))
	}

	fmt.Printf(command.HelpTopic, constants.ExeName)
}
