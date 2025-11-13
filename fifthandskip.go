package main

import "strings"

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	seq := make([]rune, 0, len(str))
	for _, r := range str {
		if r != ' ' {
			seq = append(seq, r)
		}
	}
	if len(seq) < 5 {
		return "Invalid Input\n"
	}
	filtered := make([]rune, 0, len(seq))
	for i, r := range seq {
		if (i+1)%6 == 0 {
			continue
		}
		filtered = append(filtered, r)
	}
	var b strings.Builder
	for i, r := range filtered {
		if i > 0 && i%5 == 0 {
			b.WriteByte(' ')
		}
		b.WriteRune(r)
	}
	b.WriteByte('\n')
	return b.String()
}
