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

	TopicNotFound = "topic %s not found.\n"
)
