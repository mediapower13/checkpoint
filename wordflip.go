package main

import (
	"strings"
)

func WordFlip(str string) string {
	// Trim spaces and split by fields to ignore multiple spaces
	trimmed := strings.TrimSpace(str)
	if trimmed == "" {
		// If the original string is empty or only spaces, return "Invalid Output" per spec
		return "Invalid Output\n"
	}
	words := strings.Fields(trimmed)
	// reverse words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ") + "\n"
}
