package main

import (
	"fmt"
	"os"
)

// Program: prints characters that appear in both arguments, without duplicates,
// in the order they appear in the first argument.
func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])

	printed := make(map[rune]bool)
	for _, r := range s1 {
		if printed[r] {
			continue
		}
		found := false
		for _, r2 := range s2 {
			if r == r2 {
				found = true
				break
			}
		}
		if found {
			fmt.Printf("%c", r)
			printed[r] = true
		}
	}
	fmt.Println()
}
