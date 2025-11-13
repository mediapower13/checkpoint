package main

func LastWord(s string) string {
	r := []rune(s)
	i := len(r) - 1
	for i >= 0 && r[i] == ' ' {
		i--
	}
	if i < 0 {
		return "\n"
	}
	j := i
	for j >= 0 && r[j] != ' ' {
		j--
	}
	return string(r[j+1:i+1]) + "\n"
}
