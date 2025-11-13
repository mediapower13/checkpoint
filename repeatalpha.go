package main

import "strings"

func RepeatAlpha(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			count := int(r-'a') + 1
			for i := 0; i < count; i++ {
				b.WriteRune(r)
			}
		} else if r >= 'A' && r <= 'Z' {
			count := int(r-'A') + 1
			for i := 0; i < count; i++ {
				b.WriteRune(r)
			}
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
