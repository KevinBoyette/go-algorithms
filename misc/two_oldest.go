package misc

import "sort"

// TwoOldestAges returns the two integer ages that are greatest.
// This is another CodeWars problem.
func TwoOldestAges(ages []int) [2]int {
	// TODO: Use the QuickSelect algorithm
	sort.Ints(ages)
	i := len(ages) - 1
	return [2]int{ages[i], ages[i-1]}
}
