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
}

type ConfigFile struct {
	filePath string
	data     struct {
		Server      string             `json:"server"`
		Languages   []subteez.Language `json:"languages"`
		Interactive bool               `json:"interactive"`
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
	return encoder.Decode(&c.data)
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

func (c *ConfigFile) AddLanguageFilter(lang subteez.Language) error {
	for _, language := range c.data.Languages {
		if language == lang {
			return errors.ErrDuplicateLanguage(lang)
		}
	}
	c.data.Languages = append(c.data.Languages, lang)
	return nil
}

func (c *ConfigFile) RemoveLanguageFilter(lang subteez.Language) error {
	length := len(c.data.Languages)
	for i := 0; i < length; i++ {
		if c.data.Languages[i] == lang {
			c.data.Languages[i] = c.data.Languages[length-1]
			c.data.Languages = c.data.Languages[:length-1]
			return nil
		}
	}
	return errors.ErrLanguageNotFound(lang)
}

func (c *ConfigFile) IsInteractive() bool {
	return c.data.Interactive
}

func (c *ConfigFile) SetInteractive(interactive bool) {
	c.data.Interactive = interactive
}
