package messages

const (
	DumpConfigHumanReadable = `Server      = %s
TUI Mode    = %t
Script Mode = %t

Enabled Language Filters:
`

	DumpConfig = `server=%s
tui=%t
script=%t
filters=%v`

	ConfigOptionNotFound = `config option "%s" not found`

	DuplicateLanguage = `language "%s" already has been added to filters`

	LanguageNotFound = `language "%s" not found`

	ConfigFileSaved = `config file saved`
)
