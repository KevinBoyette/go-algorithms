package sorts_test

import "testing"

type testCase struct {
	testParam []int
	expected  []int
}

// Cases generates all the test cases for the sorting functions
func testTable() map[string]testCase {
	cases := map[string]testCase{
		"empty":                                 {[]int{}, []int{}},
		"already sorted":                        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"single element":                        {[]int{1}, []int{1}},
		"negative elements":                     {[]int{-1, -2, 2, 0, 3}, []int{-2, -1, 0, 2, 3}},
		"reversed slice":                        {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"multiple versions of similar elements": {[]int{5, 4, 3, 4, 1, 2, 1, 3, 4, 5}, []int{1, 1, 2, 3, 3, 4, 4, 4, 5, 5}},
		"all the same element":                  {[]int{1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1}},
	}
	return cases
}

func runTests(f func([]int) []int, testTable map[string]testCase, t *testing.T) {
	for name, tc := range testTable {
		t.Run(name, func(t *testing.T) {
			actual := f(tc.testParam)
			expected := tc.expected
			testCase := true
			for index := range actual {
				// This if can result in panic: index out of range
				if actual[index] != expected[index] {
					testCase = false
				}
			}
			if testCase != true {
				t.Errorf("During %s; expected %v and got %v",
					name,
					tc.expected,
					actual,
				)
			}
		})
	}

}
