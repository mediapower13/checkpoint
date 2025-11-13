package main

func ConcatAlternate(slice1, slice2 []int) []int {
	len1, len2 := len(slice1), len(slice2)
	if len1 == 0 && len2 == 0 {
		return []int{}
	}
	out := make([]int, 0, len1+len2)
	turnSlice1 := len1 >= len2 // start with largest, if equal start with first
	i, j := 0, 0
	for i < len1 || j < len2 {
		if turnSlice1 {
			if i < len1 {
				out = append(out, slice1[i])
				i++
			} else if j < len2 {
				out = append(out, slice2[j])
				j++
			}
		} else {
			if j < len2 {
				out = append(out, slice2[j])
				j++
			} else if i < len1 {
				out = append(out, slice1[i])
				i++
			}
		}
		turnSlice1 = !turnSlice1
	}
	return out
}
