package main

import (
	"strings"
	"unicode"
)

func IsCapitalized(s string) bool {
	words := strings.Fields(s)
	if len(words) == 0 {
		return false
	}
	for _, w := range words {
		r := []rune(w)[0]
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			return false
		}
	}
	return true
}
