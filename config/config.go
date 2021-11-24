package config

import (
	"encoding/json"
	"os"
	"subteez/subteez"
)

type Config interface {
	GetServer() string
	SetServer(string)

	GetLanguageFilters() []subteez.Language
	SetLanguageFilters([]subteez.Language)
	ClearLanguageFilters()
	AddLanguageFilter(subteez.Language)
	RemoveLanguageFilter(subteez.Language)
}

type ConfigFile struct {
	filePath string
	data     struct {
		Server    string             `json:"server"`
		Languages []subteez.Language `json:"languages"`
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

func (c *ConfigFile) AddLanguageFilter(lang subteez.Language) {
	c.data.Languages = append(c.data.Languages, lang)
}

func (c *ConfigFile) RemoveLanguageFilter(lang subteez.Language) {
	length := len(c.data.Languages)
	for i := 0; i < length; i++ {
		if c.data.Languages[i] == lang {
			c.data.Languages[i] = c.data.Languages[length-1]
			c.data.Languages = c.data.Languages[:length-1]
			return
		}
	}
}
