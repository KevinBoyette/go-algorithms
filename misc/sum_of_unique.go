package misc

// SumOfUniqueElements returns the sum of the elements
// by skipping duplicates
//
// https://leetcode.com/problems/sum-of-unique-elements/
func SumOfUniqueElements(nums []int) int {
	const max = 101
	results := [max]int{}
	for _, i := range nums {
		results[i]++
	}

	total := 0
	for i := 1; i < max; i++ {
		if results[i] == 1 {
			total += i
		}
	}

	return total
}
