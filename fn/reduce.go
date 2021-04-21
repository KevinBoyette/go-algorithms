package fn

import "fmt"

// Reduce is not general. This function represents the
// idea behind using reduce.
func Reduce(f func(int) int, list []int) int {
	var returnValue int
	for index, listValue := range list {
		if index == 0 {
			returnValue = listValue
		} else {
			returnValue += f(listValue)

		}
		fmt.Println(returnValue)
	}
	return returnValue
}
