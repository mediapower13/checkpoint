package main

import (
	"fmt"
	"os"
)

func isBalanced(s string) bool {
	stack := make([]rune, 0, len(s))
	hasBracket := false
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			hasBracket = true
			stack = append(stack, r)
		case ')':
			hasBracket = true
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			hasBracket = true
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			hasBracket = true
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			// ignore other chars
		}
	}
	// If there were no brackets, it's considered correctly bracketed
	if !hasBracket {
		return true
	}
	return len(stack) == 0
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	for _, a := range args {
		if isBalanced(a) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
