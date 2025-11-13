package main

import "strings"

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	r := []rune(arg)
	var b strings.Builder
	save := true
	for i := 0; i < len(r); i += num {
		end := i + num
		if end > len(r) {
			end = len(r)
		}
		if save {
			for _, ch := range r[i:end] {
				b.WriteRune(ch)
			}
		}
		save = !save
	}
	return b.String()
}
