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
