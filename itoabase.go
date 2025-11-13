package main

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return "" // as per spec, only valid inputs will be tested
	}
	if value == 0 {
		return "0"
	}
	digits := "0123456789ABCDEF"
	neg := value < 0
	var v int
	if neg {
		v = -value
	} else {
		v = value
	}
	b := make([]byte, 0, 32)
	for v > 0 {
		d := v % base
		b = append(b, digits[d])
		v /= base
	}
	if neg {
		b = append(b, '-')
	}
	// reverse b
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
