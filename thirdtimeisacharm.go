package main

func ThirdTimeIsACharm(str string) string {
	r := []rune(str)
	if len(r) < 3 {
		return "\n"
	}
	out := make([]rune, 0, len(r)/3)
	for i := 2; i < len(r); i += 3 {
		out = append(out, r[i])
	}
	return string(out) + "\n"
}
