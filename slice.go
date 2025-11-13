package main

func Slice(a []string, nbrs ...int) []string {
	n := len(a)
	if len(nbrs) == 0 {
		return nil
	}
	norm := func(idx int) int {
		if idx < 0 {
			idx = n + idx
		}
		return idx
	}
	start := norm(nbrs[0])
	if start < 0 || start > n {
		return nil
	}
	if len(nbrs) == 1 {
		if start == n {
			return []string{}
		}
		return a[start:]
	}
	end := norm(nbrs[1])
	if end < 0 || end > n {
		return nil
	}
	if start >= end {
		return nil
	}
	return a[start:end]
}
