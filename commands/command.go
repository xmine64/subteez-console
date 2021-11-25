package commands

import "subteez/config"

type Command struct {
	Name        string
	Description string
	Main        func(args []string, cfg config.Config) error
}

var AllCommands = map[string]Command{
	"search": {
		Name:        "search",
		Description: "Search for movie or series titles",
		Main:        mainSearch,
	},
	"files": {
		Name:        "files",
		Description: "List all available subtitle files",
		Main:        mainFiles,
	},
	"download": {
		Name:        "download",
		Description: "Download a subtitle file",
		Main:        mainDownload,
	},
}
