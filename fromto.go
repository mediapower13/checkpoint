package main

import (
	"fmt"
	"strings"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}
	var b strings.Builder
	if from <= to {
		for i := from; i <= to; i++ {
			if i > from {
				b.WriteString(", ")
			}
			b.WriteString(fmt.Sprintf("%02d", i))
		}
	} else {
		for i := from; i >= to; i-- {
			if i < from {
				b.WriteString(", ")
			}
			b.WriteString(fmt.Sprintf("%02d", i))
		}
	}
	b.WriteByte('\n')
	return b.String()
}
