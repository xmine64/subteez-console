package config

import (
	"encoding/json"
	"os"
	"subteez/errors"
	"subteez/subteez"
)

type Config interface {
	GetServer() string
	SetServer(string)

	GetLanguageFilters() []subteez.Language
	SetLanguageFilters([]subteez.Language)
	ClearLanguageFilters()
	AddLanguageFilter(subteez.Language) error
	RemoveLanguageFilter(subteez.Language) error
	IsInteractive() bool
	SetInteractive(bool)
	IsScriptMode() bool
	SetScriptMode(bool)
}

type ConfigFile struct {
	filePath string
	data     struct {
		Server      string             `json:"server"`
		Languages   []subteez.Language `json:"languages"`
		Interactive bool               `json:"tui"`
		ScriptMode  bool               `json:"script_mode"`
	}
}

func NewConfigFile(filePath string) *ConfigFile {
	result := new(ConfigFile)
	result.filePath = filePath
	return result
}

func (c *ConfigFile) Save() error {
	file, err := os.Create(c.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(c.data)
}

func (c *ConfigFile) Load() error {
	file, err := os.Open(c.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewDecoder(file)
	if err := encoder.Decode(&c.data); err != nil {
		return err
	}
	// interactive mode and script mode can't be enabled at same time
	if c.data.Interactive && c.data.ScriptMode {
		return errors.ErrInteractiveAndScript
	}
	return nil
}

func (c *ConfigFile) GetServer() string {
	return c.data.Server
}

func (c *ConfigFile) SetServer(server string) {
	c.data.Server = server
}

func (c *ConfigFile) GetLanguageFilters() []subteez.Language {
	return c.data.Languages
}

func (c *ConfigFile) SetLanguageFilters(filters []subteez.Language) {
	c.data.Languages = filters
}

func (c *ConfigFile) ClearLanguageFilters() {
	c.data.Languages = c.data.Languages[:0]
}

func (c *ConfigFile) AddLanguageFilter(value subteez.Language) error {
	// return error if value is already in the language filter list
	for _, language := range c.data.Languages {
		if language == value {
			return errors.ErrDuplicateLanguage(value)
		}
	}

	c.data.Languages = append(c.data.Languages, value)
	return nil
}

func (c *ConfigFile) RemoveLanguageFilter(value subteez.Language) error {
	// find value in list and remove it
	length := len(c.data.Languages)
	for i := 0; i < length; i++ {
		if c.data.Languages[i] == value {
			// shift items one cell back, and then remove last cell
			for j := i + 1; j < length; j++ {
				c.data.Languages[j-1] = c.data.Languages[j]
			}
			c.data.Languages = c.data.Languages[:length-1]
			return nil
		}
	}

	// error if value not found
	return errors.ErrLanguageNotFound(value)
}

func (c *ConfigFile) IsInteractive() bool {
	return c.data.Interactive
}

func (c *ConfigFile) SetInteractive(interactive bool) {
	c.data.Interactive = interactive
}

func (c *ConfigFile) IsScriptMode() bool {
	return c.data.ScriptMode
}

func (c *ConfigFile) SetScriptMode(value bool) {
	c.data.ScriptMode = value
}
