package sorts

import "math/rand"

// QuickSortRecur will sort your slice quickly, but also
// recursively. Should not be used with large slices
func QuickSortRecur(list []int) []int {
	listLength := len(list)
	if listLength <= 1 {
		return list
	}
	left, right := 0, listLength-1
	pivotPoint := rand.Int() % listLength
	list[pivotPoint], list[right] = list[right], list[pivotPoint]

	for index := range list {
		if list[index] < list[right] {
			list[index], list[left] = list[left], list[index]
			left++
		}
	}

	list[left], list[right] = list[right], list[left]
	QuickSortRecur(list[left+1:])
	QuickSortRecur(list[:left])
	return list
}
