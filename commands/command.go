package commands

import (
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
		Main:        mainSearch,
	},
	"files": {
		Name:        "files",
		Description: messages.FilesDescription,
		Main:        mainFiles,
	},
	"download": {
		Name:        "download",
		Description: messages.DownloadDescription,
		Main:        mainDownload,
	},
	"config": {
		Name:        "config",
		Description: messages.ConfigDescription,
		Main:        mainConfig,
	},
}
