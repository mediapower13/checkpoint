package main

func FirstWord(s string) string {
	i := 0
	n := len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	if i >= n {
		return "\n"
	}
	j := i
	for j < n && s[j] != ' ' {
		j++
	}
	return s[i:j] + "\n"
}
