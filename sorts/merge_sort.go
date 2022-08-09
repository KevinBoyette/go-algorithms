package sorts

// MergeSort takes in a slice of ints and
// returns a sorted slice of ints
func MergeSort(list []int) []int {
	listLength := len(list)
	if listLength <= 1 {
		return list
	}
	mid := listLength / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return Merge(left, right)

}

// Merge is a helper function for MergeSort
//
// Example:
//
//	Merge([]int{2,0,3},[]int{-1,0 -2}) -> []int{-1, 0, -2, 2, 0, 3}
func Merge(left, right []int) []int {
	ret := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(ret, right...)
		}
		if len(right) == 0 {
			return append(ret, left...)
		}
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}
	return ret
}
