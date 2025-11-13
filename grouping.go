package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	pattern := os.Args[1]
	text := os.Args[2]

	if pattern == "" || text == "" {
		return
	}

	// pattern should contain parentheses; we'll accept simple patterns like (a) or (e|n)
	// Validate by compiling as regexp
	re, err := regexp.Compile(pattern)
	if err != nil {
		return
	}

	// Find all matches; we want the whole word that contains the match.
	// We'll scan through matches and expand to word boundaries.
	matches := re.FindAllStringIndex(text, -1)
	if len(matches) == 0 {
		return
	}

	results := []string{}
	for _, m := range matches {
		start, end := m[0], m[1]
		// expand start backwards to word boundary
		s := start
		for s > 0 {
			r := text[s-1]
			if !(isWordChar(r)) {
				break
			}
			s--
		}
		e := end
		for e < len(text) {
			r := text[e]
			if !(isWordChar(r)) {
				break
			}
			e++
		}
		word := text[s:e]
		word = strings.TrimSpace(word)
		if word != "" {
			results = append(results, word)
		}
	}

	if len(results) == 0 {
		return
	}

	for i, w := range results {
		fmt.Printf("%d: %s\n", i+1, w)
	}
}

func isWordChar(b byte) bool {
	// consider letters, digits and apostrophes as part of words
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '\'' {
		return true
	}
	return false
}
