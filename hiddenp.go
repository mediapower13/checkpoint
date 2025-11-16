package main

import (
	"fmt"
	"os"
)

func mainh() {
	if len(os.Args) != 3 {
		return
	}
	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])

	if len(s1) == 0 {
		fmt.Println("1")
		return
	}

	j := 0
	for _, r := range s2 {
		if r == s1[j] {
			j++
			if j == len(s1) {
				fmt.Println("1")
				return
			}
		}
	}
	fmt.Println("0")
}
