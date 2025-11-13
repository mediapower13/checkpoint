package main

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return []int{}
	}
	out := make([]int, 0, len(slice1)+len(slice2))
	out = append(out, slice1...)
	out = append(out, slice2...)
	return out
}
