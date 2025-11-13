package main

import (
	"fmt"
	"os"
	"unicode"
)

func transform(s string) string {
	r := []rune(s)
	out := make([]rune, 0, len(r))
	n := len(r)
	i := 0
	for i < n {
		if r[i] == ' ' {
			out = append(out, r[i])
			i++
			continue
		}
		j := i
		for j < n && r[j] != ' ' {
			j++
		}
		if j-i == 1 {
			out = append(out, unicode.ToUpper(r[i]))
		} else if j-i > 1 {
			for k := i; k < j-1; k++ {
				out = append(out, unicode.ToLower(r[k]))
			}
			out = append(out, unicode.ToUpper(r[j-1]))
		}
		i = j
	}
	return string(out)
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	for _, arg := range os.Args[1:] {
		fmt.Println(transform(arg))
	}
}
