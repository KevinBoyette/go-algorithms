package misc

import (
	"math"
)

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
	maxSoFar := math.MinInt32
	current := math.MinInt32
	for _, i := range nums {
		if current < 0 {
			current = i
		} else {
			current += i
		}
		maxSoFar = maxOf(current, maxSoFar)
	}
	return maxSoFar
}
