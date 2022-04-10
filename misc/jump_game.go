package misc

// max computes the maximum of two integers
func max(first, next int) int {
	if first >= next {
		return first
	}
	return next
}

// CanJump solves the problem defined at
// https://leetcode.com/problems/jump-game
func CanJump(nums []int) bool {
	last := 0
	numsLength := len(nums) - 1
	for i := 0; i < numsLength; i++ {
		if i > last {
			break
		}
		last = max(last, nums[i]+i)
	}
	return last >= numsLength
}
