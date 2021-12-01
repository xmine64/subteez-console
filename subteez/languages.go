package subteez

import (
	"strconv"
	"strings"
)

var Languages = []Language{
	"en", "fa", "ar", "hi", "de", "fr", "it", "pl", "ru", "es", "tr",
}

func (lang Language) GetTitle() string {
	titles := map[Language]string{
		"en": "English",
		"fa": "Persian",
		"ar": "Arabic",
		"hi": "Hindi",
		"de": "German",
		"fr": "French",
		"it": "Italian",
		"pl": "Polish",
		"ru": "Russian",
		"es": "Spanish",
		"tr": "Turkish",
	}
	result, exists := titles[lang]
	if !exists {
		return "Unknown"
	}
	return result
}

func (lang Language) GetDownloadPathPart() (string, error) {
	id := map[Language]string{
		"en": "english",
		"fa": "farsi_persian",
		"ar": "arabic",
		"hi": "hindi",
		"de": "german",
		"fr": "french",
		"it": "italian",
		"pl": "polish",
		"ru": "russian",
		"es": "spanish",
		"tr": "turkish",
	}
	result, exists := id[lang]
	if !exists {
		return "", strconv.ErrRange
	}
	return result, nil
}

func ParseLanguage(str string) (Language, error) {
	for _, language := range Languages {
		if strings.EqualFold(string(language), str) {
			return language, nil
		}
		if strings.EqualFold(language.GetTitle(), str) {
			return language, nil
		}
	}
	return "", strconv.ErrRange
}
