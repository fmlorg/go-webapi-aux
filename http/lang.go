package http

import (
	"strings"

	"golang.org/x/text/language"
)

func languageTagSelect(list ...string) string {
	// languages supported by this service
	matcher := language.NewMatcher([]language.Tag{
		language.English, language.Japanese,
	})

	tag, _ := language.MatchStrings(matcher, list...)
	return tag.String()
}

func LanguageSelect(list ...string) string {
	tag := languageTagSelect(list...)

	r := "en"
	if strings.HasPrefix(tag, "ja") {
		r = "ja"
	} else {
		r = "en"
	}

	return r
}
