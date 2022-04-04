package misc

import (
	"testing"
)

func TestTwoOldestAges(t *testing.T) {
	cases := map[string]struct {
		param    []int
		expected [2]int
	}{
		"simplest case":      {[]int{1, 2}, [2]int{1, 2}},
		"already sorted":     {[]int{1, 2, 3, 5}, [2]int{3, 5}},
		"reversed":           {[]int{5, 4, 3, 2, 1}, [2]int{4, 5}},
		"codewars example 1": {[]int{6, 5, 83, 5, 3, 18}, [2]int{18, 83}},
		"codewars example 2": {[]int{1, 5, 87, 45, 8, 8}, [2]int{45, 87}},
	}
	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			actual := TwoOldestAges(testCase.param)
			expected := testCase.expected
			if actual != expected {
				t.Errorf("During %s; expected %v and got %v",
					name,
					testCase.expected,
					actual,
				)
			}
		})
	}
}
