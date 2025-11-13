package main

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	var t int64 = int64(n)
	if t < 0 {
		t = -t
	}
	b := make([]byte, 0, 20)
	for t > 0 {
		d := byte(t%10) + '0'
		b = append(b, d)
		t /= 10
	}
	if neg {
		b = append(b, '-')
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
