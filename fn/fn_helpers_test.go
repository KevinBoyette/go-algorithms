package fn

type mapTestCase struct {
	testFunc  func(int) int
	testSlice []int
	expected  []int
}
type reduceTestCase struct {
	testFunc  func(int) int
	testSlice []int
	expected  int
}

func mapTestTable() map[string]mapTestCase {
	return map[string]mapTestCase{
		"plus1":          {plus1, []int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 6}},
		"plus2":          {plus2, []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}},
		"plus3":          {plus3, []int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}},
		"plus10":         {plus10, []int{1, 2, 3, 4, 5}, []int{11, 12, 13, 14, 15}},
		"single element": {plus10, []int{1}, []int{11}},
		"empty slice":    {plus10, []int{}, []int{}},
	}
}

func reduceTestTable() map[string]reduceTestCase {
	return map[string]reduceTestCase{
		// viewed as reduce with sum of (x + y + 1)
		"plus1": {plus1, []int{1, 2, 3, 4, 5}, 19},
		// viewed as reduce with sum of (x + y + 2)
		"plus2":                         {plus2, []int{1, 2, 3, 4, 5}, 23},
		"plus3":                         {plus3, []int{1, 2, 3, 4, 5}, 27},
		"plus10":                        {plus10, []int{1, 2, 3, 4, 5}, 55},
		"single element":                {plus10, []int{1}, 1},
		"empty slice":                   {plus10, []int{}, 0},
		"reduce example: plus1":         {plus1, []int{1, 3, 4, 10, 4}, 26},
		"reduce example(sum): identity": {identity, []int{1, 3, 4, 10, 4}, 22},
		"(sum 1,2,3,4,5): identity":     {identity, []int{1, 2, 3, 4, 5}, 15},
	}
}

func identity(x int) int {
	return x
}
func plus1(x int) int {
	return x + 1
}

func plus2(x int) int {
	return x + 2
}

func plus3(x int) int {
	return x + 3
}

func plus10(x int) int {
	return x + 10
}
