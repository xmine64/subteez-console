// commands definition

package cli

import (
	"subteez/cli/config_command"
	"subteez/cli/download_command"
	"subteez/cli/files_command"
	"subteez/cli/search_command"
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
		Main:        search_command.Main,
	},
	"files": {
		Name:        "files",
		Description: messages.FilesDescription,
		HelpTopic:   messages.FilesHelpTopic,
		Main:        files_command.Main,
	},
	"download": {
		Name:        "download",
		Description: messages.DownloadDescription,
		HelpTopic:   messages.DownloadHelpTopic,
		Main:        download_command.Main,
	},
	"config": {
		Name:        "config",
		Description: messages.ConfigDescription,
		HelpTopic:   messages.ConfigHelpTopic,
		Main:        config_command.Main,
	},
}
