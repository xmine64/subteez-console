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
	Main        func(args []string, cfg config.Config) error
}

var AllCommands = map[string]Command{
	"search": {
		Name:        "search",
		Description: messages.SearchDescription,
		Main:        search.Main,
	},
	"files": {
		Name:        "files",
		Description: messages.FilesDescription,
		Main:        files.Main,
	},
	"download": {
		Name:        "download",
		Description: messages.DownloadDescription,
		Main:        download.Main,
	},
	"config": {
		Name:        "config",
		Description: messages.ConfigDescription,
		Main:        configCommand.Main,
	},
}
