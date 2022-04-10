package misc

// https://leetcode.com/problems/boats-to-save-people/
import "sort"

// NumRescueBoats needed to save people.
func NumRescueBoats(people []int, limit int) int {
	numberOfTrips := len(people)
	i := 0
	j := numberOfTrips - 1
	sort.Ints(people)
	for i < j {
		if people[i]+people[j] <= limit {
			numberOfTrips--
			i++
		}
		j--
	}

	return numberOfTrips
}
