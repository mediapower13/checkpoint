package main

func HashCode(dec string) string {
	if dec == "" {
		return ""
	}
	n := len(dec)
	out := make([]rune, 0, len(dec))
	for _, r := range dec {
		v := (int(r) + n) % 127
		if v < 32 {
			v += 33
		}
		out = append(out, rune(v))
	}
	return string(out)
}
