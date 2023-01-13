package regex

import (
	"regexp"
	"strings"
	"unicode"
)

//Checks if string contains a uppercase letter
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

//If string is uppercase, separate string to hyphens
func StringToHyphen(word string) string {
	if IsUpper(word) {
		m1 := regexp.MustCompile(`([a-z])([A-Z])`)

		return strings.ToLower(m1.ReplaceAllString(word, "$1-$2"))
	}

	return strings.ToLower(word)
}
