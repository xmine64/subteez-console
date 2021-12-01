package subteez

import (
	"strconv"
	"strings"
)

type Language string

// valid values of language
var (
	EnglishLanguage = Language("en")
	PersianLanguage = Language("fa")
	ArabicLanguage  = Language("ar")
	HindiLanguage   = Language("hi")
	GermanLanguage  = Language("de")
	FrenchLanguage  = Language("fr")
	ItalianLanguage = Language("it")
	PolishLanguage  = Language("pl")
	RussianLanguage = Language("ru")
	SpanishLanguage = Language("es")
	TurkishLanguage = Language("tr")

	Languages = []Language{
		EnglishLanguage,
		PersianLanguage,
		ArabicLanguage,
		HindiLanguage,
		GermanLanguage,
		FrenchLanguage,
		ItalianLanguage,
		PolishLanguage,
		RussianLanguage,
		SpanishLanguage,
		TurkishLanguage,
	}
)

// get display name of language
func (value Language) GetTitle() string {
	titles := map[Language]string{
		EnglishLanguage: "English",
		PersianLanguage: "Persian",
		ArabicLanguage:  "Arabic",
		HindiLanguage:   "Hindi",
		GermanLanguage:  "German",
		FrenchLanguage:  "French",
		ItalianLanguage: "Italian",
		PolishLanguage:  "Polish",
		RussianLanguage: "Russian",
		SpanishLanguage: "Spanish",
		TurkishLanguage: "Turkish",
	}
	if result, exists := titles[value]; exists {
		return result
	} else {
		return "Uknown"
	}
}

// get language code
func (value Language) GetCode() string {
	return string(value)
}

// get language code used for generating download links
func (value Language) GetDownloadPathPart() (string, error) {
	id := map[Language]string{
		EnglishLanguage: "english",
		PersianLanguage: "farsi_persian",
		ArabicLanguage:  "arabic",
		HindiLanguage:   "hindi",
		GermanLanguage:  "german",
		FrenchLanguage:  "french",
		ItalianLanguage: "italian",
		PolishLanguage:  "polish",
		RussianLanguage: "russian",
		SpanishLanguage: "spanish",
		TurkishLanguage: "turkish",
	}
	if result, exists := id[value]; exists {
		return result, nil
	} else {
		return "", strconv.ErrRange
	}
}

// parse string to Language
func ParseLanguage(str string) (Language, error) {
	for _, language := range Languages {
		if strings.EqualFold(language.GetCode(), str) || strings.EqualFold(language.GetTitle(), str) {
			return language, nil
		}
	}
	return "", strconv.ErrRange
}
