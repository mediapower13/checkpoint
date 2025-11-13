package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	s := os.Args[1]
	fields := strings.Fields(s)
	if len(fields) == 0 {
		fmt.Println()
		return
	}
	if len(fields) == 1 {
		fmt.Println(fields[0])
		return
	}
	rot := append(fields[1:], fields[0])
	fmt.Println(strings.Join(rot, " "))
}
