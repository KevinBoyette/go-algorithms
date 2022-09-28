package sorts

// BubbleSort will sort a slice of ints
//
// You probably shouldn't use BubbleSort in a
// real program
func BubbleSort(list []int) []int {
	listLength := len(list)
	if listLength <= 0 {
		return list
	}
	for i := 1; i < listLength; i++ {
		for j := 0; j < listLength-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}

	return list
}
