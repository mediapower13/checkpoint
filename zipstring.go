package main

import (
	"strconv"
	"strings"
)

func ZipString(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	var b strings.Builder
	count := 1
	for i := 1; i <= len(r); i++ {
		if i < len(r) && r[i] == r[i-1] {
			count++
			continue
		}
		b.WriteString(strconv.Itoa(count))
		b.WriteRune(r[i-1])
		count = 1
	}
	return b.String()
}
