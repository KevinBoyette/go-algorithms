package fn

// Map calls a function for each element in
// a slice
func Map(f func(int) int, list []int) []int {
	temp := make([]int, len(list))
	for index, elem := range list {
		temp[index] = f(elem)
	}
	return temp
}
