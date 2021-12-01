package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"subteez/cli"
	"subteez/config"
	"subteez/constants"
	"subteez/errors"
	"subteez/messages"
	"subteez/subteez"
	"subteez/tui"
)

func main() {
	// find config folder
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	configDir = path.Join(configDir, constants.VendorName, constants.Name)
	if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	configFilePath := path.Join(configDir, constants.ConfigFileName)

	// set command-line flags
	configFileFlag := flag.String("config", configFilePath, "")
	interactiveFlag := flag.Bool("interactive", false, "")
	scriptFlag := flag.Bool("script", false, "")
	helpFlag := flag.Bool("help", false, "")

	flag.Parse()

	// load config file, and if it doesn't exist create a default one
	configFile := config.NewConfigFile(*configFileFlag)
	if configFile.Load() != nil {
		configFile.SetServer(constants.DefaultServer)
		configFile.SetLanguageFilters(subteez.Languages)
		err := configFile.Save()
		if err != nil {
			log.Fatal(err)
		}
	}

	// if stdout is not terminal, enable script mode automatically
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		configFile.SetInteractive(false)
		configFile.SetScriptMode(true)
	}

	if *helpFlag {
		showHelpTopic(flag.Arg(0))
		return
	}

	// interactive flag and script flag can't be used togother
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
		log.Printf(messages.Version,
			constants.Name, constants.VersionMajor, constants.VersionMinor, constants.VersionBuild, constants.VendorName)
	}

	// run interactive mode if it's enabled and no command is given
	if flag.NArg() < 1 {
		if configFile.IsInteractive() {
			context := &tui.Context{}
			context.Initialize(configFile)
			if err := context.Run(); err != nil {
				log.Fatal(err)
			}
			return
		} else {
			showHelp()
			log.Fatal(errors.ErrNotEnoughArguments)
		}
	}

	// run given command if it exists, else show help message
	if command, exists := cli.AllCommands[flag.Arg(0)]; exists {
		if err := command.Main(flag.Args(), configFile); err != nil {
			// save config file if errors.ErrConfigChanged returned
			if err == errors.ErrConfigChanged {
				configFile.Save()
				log.Print(messages.ConfigFileSaved)
				return
			}

			log.Fatal(err)
		}
		if !configFile.IsScriptMode() {
			log.Println(messages.Done)
		}
	} else {
		if !configFile.IsScriptMode() {
			showHelp()
		}
		log.Fatal(errors.ErrCommandNotFound(flag.Arg(0)))
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

	for _, command := range cli.AllCommands {
		fmt.Printf(messages.CommandRow, command.Name, command.Description)
	}
}

func showHelpTopic(topic string) {
	if topic == "" {
		showHelp()
		return
	}

	// print given help topic if it exists, else show main help topic
	if command, exists := cli.AllCommands[topic]; exists {
		fmt.Printf(command.HelpTopic, constants.ExeName)
	} else {
		showHelp()
		log.Fatal(errors.ErrHelpTopicNotFound(topic))
	}
}
