package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil || n <= 1 {
		return
	}
	var parts []string
	for n%2 == 0 {
		parts = append(parts, "2")
		n /= 2
	}
	for i := int64(3); i*i <= n; i += 2 {
		for n%i == 0 {
			parts = append(parts, strconv.FormatInt(i, 10))
			n /= i
		}
	}
	if n > 1 {
		parts = append(parts, strconv.FormatInt(n, 10))
	}
	if len(parts) == 0 {
		return
	}
	fmt.Println(strings.Join(parts, "*"))
}
