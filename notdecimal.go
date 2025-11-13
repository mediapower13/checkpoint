package main

import "strings"

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	orig := dec
	sign := ""
	if dec[0] == '+' || dec[0] == '-' {
		if dec[0] == '-' {
			sign = "-"
		}
		dec = dec[1:]
	}
	if dec == "" {
		return orig + "\n"
	}
	dotCount := 0
	for _, r := range dec {
		if r == '.' {
			dotCount++
			if dotCount > 1 {
				return orig + "\n"
			}
			continue
		}
		if r < '0' || r > '9' {
			return orig + "\n"
		}
	}
	if dotCount == 0 {
		return orig + "\n"
	}
	parts := strings.SplitN(dec, ".", 2)
	intPart := parts[0]
	fracPart := parts[1]
	if len(fracPart) == 0 {
		return orig + "\n"
	}
	allZero := true
	for _, r := range fracPart {
		if r != '0' {
			allZero = false
			break
		}
	}
	if allZero {
		return orig + "\n"
	}
	concat := intPart + fracPart
	i := 0
	for i < len(concat) && concat[i] == '0' {
		i++
	}
	trimmed := concat[i:]
	if trimmed == "" {
		trimmed = "0"
	}
	if sign == "-" {
		trimmed = "-" + trimmed
	}
	return trimmed + "\n"
}
