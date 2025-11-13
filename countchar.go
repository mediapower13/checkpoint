package main

func CountChar(str string, c rune) int {
	if len(str) == 0 {
		return 0
	}
	count := 0
	for _, r := range str {
		if r == c {
			count++
		}
	}
	return count
}
