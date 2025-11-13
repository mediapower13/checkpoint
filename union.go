package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	a := []rune(os.Args[1])
	b := []rune(os.Args[2])
	seen := make(map[rune]bool)
	for _, r := range a {
		if !seen[r] {
			fmt.Printf("%c", r)
			seen[r] = true
		}
	}
	for _, r := range b {
		if !seen[r] {
			fmt.Printf("%c", r)
			seen[r] = true
		}
	}
	fmt.Println()
}
