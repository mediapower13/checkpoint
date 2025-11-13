package main

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	n := len(runes)

	// validate: only letters, no trailing uppercase, no two consecutive uppercase
	for i, r := range runes {
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			return s
		}
		if i > 0 {
			prev := runes[i-1]
			if (r >= 'A' && r <= 'Z') && (prev >= 'A' && prev <= 'Z') {
				return s
			}
		}
	}
	if runes[n-1] >= 'A' && runes[n-1] <= 'Z' {
		return s
	}

	out := make([]rune, 0, n*2)
	for i, r := range runes {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out = append(out, '_')
		}
		out = append(out, r)
	}
	return string(out)
}
