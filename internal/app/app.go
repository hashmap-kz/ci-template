package app

import (
	"unicode"
	"unicode/utf8"
)

// IsValidIdentifier checks if the given string is a valid Go identifier.
func IsValidIdentifier(s string) bool {
	if s == "" {
		return false
	}

	// First character must be a letter or underscore
	firstRune, _ := utf8.DecodeRuneInString(s)
	if !unicode.IsLetter(firstRune) && firstRune != '_' {
		return false
	}

	// Remaining characters can be letters, digits, or underscores
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}

	return true
}
