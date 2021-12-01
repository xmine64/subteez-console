package messages

const (
	HelpMessage = `%s v%d.%d.%d

Usage: %s [-help | -interactive] <Command> <Arguments>
	
Options:
    
    -help [command] : Show help message

    -tui            : Enable interactive TUI mode

    -script         : Enable script mode
                      Script mode disables log and formats output for using in scripts

Commands:
	
`
	CommandRow = "    %s\t: %s\n\n"

	SearchHelpTopic = `Usage: %s search <Query>

Search movies or series for given query and print the result or show a menu on TUI mode.

`
	FilesHelpTopic = `Usage: %s files <Movie ID>
	
Show available files for given movie or series ID and print the result or show a menu on TUI mode.

Example: files arcane-league-of-legends-first-season

Example using full address:
         files /subtitles/arcane-league-of-legends-first-season 

`
	DownloadHelpTopic = `Usage: %s download <Movie ID> <Language> <File ID>
	
Download file with specified file ID and save it in a file, or write output in stdout on script mode.

Examples: download arcane-league-of-legends-first-season fa 2625907
          download arcane-league-of-legends-first-season Persian 2625907

Example using full address:
          download /subtitles/arcane-league-of-legends-first-season/farsi_persian/2625907

`

	ConfigHelpTopic = `Usage: %s config [command]
	
Display or Change configurations, using commands or using TUI menu.

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
