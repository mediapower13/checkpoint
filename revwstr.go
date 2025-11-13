package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	if input == "" {
		// no output for empty string per spec
		return
	}
	words := strings.Split(input, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if i != len(words)-1 {
			fmt.Print(" ")
		}
		fmt.Print(words[i])
	}
	fmt.Println()
}
