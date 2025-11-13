package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
}

func main() {
	args := os.Args[1:]
	// If no args => print usage
	if len(args) == 0 {
		printUsage()
		return
	}

	// If any argument starts with '-' and has 'h' as the first option char (i.e. second rune), print usage
	for _, a := range args {
		if len(a) >= 2 && a[0] == '-' && a[1] == 'h' {
			printUsage()
			return
		}
	}

	var opts uint32 = 0

	for _, a := range args {
		if len(a) < 2 || a[0] != '-' {
			// not an option, ignore per spec (though inputs won't include these)
			continue
		}
		if a == "-" {
			fmt.Println("Invalid Option")
			return
		}
		for _, ch := range a[1:] {
			if ch < 'a' || ch > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			pos := ch - 'a' // a -> 0, b -> 1, ... z -> 25
			opts |= 1 << pos
		}
	}

	// Print 4 groups of 8 bits starting from most significant byte
	b3 := (opts >> 24) & 0xFF
	b2 := (opts >> 16) & 0xFF
	b1 := (opts >> 8) & 0xFF
	b0 := opts & 0xFF

	fmt.Printf("%08b %08b %08b %08b\n", b3, b2, b1, b0)
}
