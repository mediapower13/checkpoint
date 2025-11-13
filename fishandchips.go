package main

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	div2 := n%2 == 0
	div3 := n%3 == 0
	if div2 && div3 {
		return "fish and chips"
	}
	if div2 {
		return "fish"
	}
	if div3 {
		return "chips"
	}
	return "error: non divisible"
}
