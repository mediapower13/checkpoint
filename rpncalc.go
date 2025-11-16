package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printError() {
	fmt.Println("Error")
}

func maini() {
	if len(os.Args) != 2 {
		printError()
		return
	}
	expr := os.Args[1]
	tokens := strings.Fields(expr)
	if len(tokens) == 0 {
		printError()
		return
	}

	var stack []int

	for _, tok := range tokens {
		// operator check: single-char tokens + - * / % are operators
		if len(tok) == 1 {
			switch tok {
			case "+", "-", "*", "/", "%":
				if len(stack) < 2 {
					printError()
					return
				}
				b := stack[len(stack)-1]
				a := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				var res int
				switch tok {
				case "+":
					res = a + b
				case "-":
					res = a - b
				case "*":
					res = a * b
				case "/":
					if b == 0 {
						printError()
						return
					}
					res = a / b
				case "%":
					if b == 0 {
						printError()
						return
					}
					res = a % b
				}
				stack = append(stack, res)
				continue
			}
		}

		// otherwise try parse integer (handles negative numbers like -3)
		n, err := strconv.Atoi(tok)
		if err != nil {
			printError()
			return
		}
		stack = append(stack, n)
	}

	if len(stack) != 1 {
		printError()
		return
	}
	fmt.Println(stack[0])
}
