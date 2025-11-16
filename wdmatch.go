package main

import (
	"fmt"
	"os"
)

func mainb() {
	if len(os.Args) != 3 {
		return
	}
	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])

	j := 0
	for _, r := range s2 {
		if r == s1[j] {
			j++
			if j == len(s1) {
				fmt.Println(string(s1))
				return
			}
		}
	}
}

