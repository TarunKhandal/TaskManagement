package helpers

import (
	"regexp"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gookit/validate"
)

var RegExpMultiSpaces = regexp.MustCompile(`\s+`)
var RegExpNotInNamify = regexp.MustCompile(`[^A-Za-z0-9 #$&()\-\\'",./]`)

func init() {
	validate.AddFilter("Namify", Namify)
	validate.AddFilter("TrimMergeWhiteSpaces", TrimMergeWhiteSpaces)
}

func ToTitle(in string) string {
	lenRunes := utf8.RuneCountInString(in)
	if lenRunes == 0 {
		return ""
	}

	// First lower case all characters
	outRunes := []rune(strings.ToLower(in))

	// First character will always be upper case
	outRunes[0] = unicode.ToUpper(outRunes[0])

	// Loop from 2nd character to second last character.
	for j := 1; j < lenRunes-1; j++ {
		if slices.Contains([]string{".", "&", " "}, string(outRunes[j])) {
			// Make next character uppercase
			outRunes[j+1] = unicode.ToUpper(outRunes[j+1])
		}
	}

	return string(outRunes)
}

func TrimMergeWhiteSpaces(val interface{}) string {
	s, ok := val.(string)
	if !ok {
		panic("invalid data type")
	}
	return strings.TrimSpace(RegExpMultiSpaces.ReplaceAllString(s, " "))
}

func Namify(val interface{}) string {
	s, ok := val.(string)
	if !ok {
		panic("invalid data type")
	}
	return ToTitle(TrimMergeWhiteSpaces(RegExpNotInNamify.ReplaceAllString(s, "")))
}
