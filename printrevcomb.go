package main

import "fmt"

func mainc() {
	first := true
	for a := 9; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			for c := b - 1; c >= 0; c-- {
				if !first {
					fmt.Print(", ")
				}
				fmt.Printf("%d%d%d", a, b, c)
				first = false
			}
		}
	}
	fmt.Println()
}
