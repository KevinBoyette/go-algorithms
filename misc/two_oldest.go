package misc

// TwoOldestAges returns the two integer ages that are greatest.
// This is another CodeWars problem.
func TwoOldestAges(ages []int) [2]int {
	var oldestTwo [2]int
	for _, age := range ages {
		if age > oldestTwo[1] {
			oldestTwo[0], oldestTwo[1] = oldestTwo[1], age
		} else if age > oldestTwo[0] {
			oldestTwo[0] = age
		}
	}
	return oldestTwo
}
