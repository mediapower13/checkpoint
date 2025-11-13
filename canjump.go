package main

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	last := len(nums) - 1
	pos := 0
	for pos < last {
		step := nums[pos]
		if step == 0 {
			return false
		}
		next := pos + int(step)
		if next > last {
			return false
		}
		if next == last {
			return true
		}
		pos = next
	}
	return pos == last
}
