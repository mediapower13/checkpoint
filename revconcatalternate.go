package main

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1, len2 := len(slice1), len(slice2)
	if len1 == 0 && len2 == 0 {
		return []int{}
	}
	i, j := len1-1, len2-1
	out := make([]int, 0, len1+len2)

	if len1 > len2 {
		for k := 0; k < len1-len2; k++ {
			out = append(out, slice1[i])
			i--
		}
	} else if len2 > len1 {
		for k := 0; k < len2-len1; k++ {
			out = append(out, slice2[j])
			j--
		}
	}

	turnFirst := true
	for i >= 0 || j >= 0 {
		if turnFirst {
			if i >= 0 {
				out = append(out, slice1[i])
				i--
			} else if j >= 0 {
				out = append(out, slice2[j])
				j--
			}
		} else {
			if j >= 0 {
				out = append(out, slice2[j])
				j--
			} else if i >= 0 {
				out = append(out, slice1[i])
				i--
			}
		}
		turnFirst = !turnFirst
	}
	return out
}
