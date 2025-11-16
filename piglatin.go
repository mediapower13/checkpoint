package main

import (
	"fmt"
	"os"
	"unicode"
)

func isVowel(r rune) bool {
	switch unicode.ToLower(r) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func mainf() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	runes := []rune(s)

	// find first vowel
	firstV := -1
	for i, r := range runes {
		if isVowel(r) {
			firstV = i
			break
		}
	}

	if firstV == 0 {
		fmt.Println(s + "ay")
		return
	}
	if firstV > 0 {
		out := string(runes[firstV:]) + string(runes[:firstV]) + "ay"
		fmt.Println(out)
		return
	}

	// no vowel found
	fmt.Println("No vowels")
}
