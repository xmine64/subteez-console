package messages

const (
	HelpMessage = `%s v%d.%d.%d

Usage: %s [-help | -interactive] <Command> <Arguments>
	
Options:
    
    -help [command] : Show help message

    -interactive    : Enable interactive mode

Commands:
	
`
	CommandRow = "    %s\t: %s\n\n"

	SearchHelpTopic = `Usage: %s [-interactive] search <Query>

Search movies or series for given query and print the result.
In interactive mode you can select a title using arrow keys.
`
	FilesHelpTopic = `Usage: %s [-interactive] files <Movie ID>
	
Show available files for given movie or series ID.
In interactive mode you can select a file using arrow keys.
`
	DownloadHelpTopic = `Usage: %s download <File ID>
	
Download file with specified file ID. Interactive mode is not supported.
`

	ConfigHelpTopic = `Usage: %s config [command]
	
Display or Change configurations.

Commands:

            show                   : show current configurations

            set [config] [value]   : set config to given value

                                     possible configs: server, interactive

            add-filter [language]  : add language to language filters

                                     example: add-filter en
                                              add-filter English

            rm-filter [language]   : remove language from language filters

                                     example: rm-filter en
                                              rm-filter English

            set-filter [languages] : set language filter

                                     example: set-filter ru
                                              set-filter en fa hi
                                              set-filter English Persian Hindi
`

	TopicNotFound = `topic "%s" not found`
)
