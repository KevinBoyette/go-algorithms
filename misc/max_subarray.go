package misc

func maxOf(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// MaxSubArray solves:
// Given an integer array nums, find the contiguous subarray  (containing at
// least one number) which has the largest sum and return its sum.
//
//A subarray is a contiguous part of an array.
//https://leetcode.com/problems/maximum-subarray/
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		maximumSum = nums[0]
		numsLength = len(nums)
	)
	for i := 1; i < numsLength; i++ {
		previous := nums[i-1]
		if previous > 0 {
			nums[i] += previous
		}
		maximumSum = maxOf(nums[i], maximumSum)
	}

	return maximumSum
}
