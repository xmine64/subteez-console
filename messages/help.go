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

Example: files arcane-league-of-legends-first-season

Example using full address:
         files /subtitles/arcane-league-of-legends-first-season 

`
	DownloadHelpTopic = `Usage: %s download <Movie ID> <Language> <File ID>
	
Download file with specified file ID. Interactive mode is not supported.

Examples: download arcane-league-of-legends-first-season fa 2625907
          download arcane-league-of-legends-first-season Persian 2625907

Example using full address:
          download /subtitles/arcane-league-of-legends-first-season/farsi_persian/2625907

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
