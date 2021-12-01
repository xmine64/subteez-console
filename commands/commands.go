package commands

import (
	configCommand "subteez/commands/config"
	"subteez/commands/download"
	"subteez/commands/files"
	"subteez/commands/search"
	"subteez/config"
	"subteez/messages"
)

type Command struct {
	Name        string
	Description string
	HelpTopic   string
	Main        func(args []string, cfg config.Config) error
}

var AllCommands = map[string]Command{
	"search": {
		Name:        "search",
		Description: messages.SearchDescription,
		HelpTopic:   messages.SearchHelpTopic,
		Main:        search.Main,
	},
	"files": {
		Name:        "files",
		Description: messages.FilesDescription,
		HelpTopic:   messages.FilesHelpTopic,
		Main:        files.Main,
	},
	"download": {
		Name:        "download",
		Description: messages.DownloadDescription,
		HelpTopic:   messages.DownloadHelpTopic,
		Main:        download.Main,
	},
	"config": {
		Name:        "config",
		Description: messages.ConfigDescription,
		HelpTopic:   "",
		Main:        configCommand.Main,
	},
}
