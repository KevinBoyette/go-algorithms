package misc_test

import (
	"kevinboyette/algorithms/misc"
	"testing"
)

func TestTwoOldestAges(t *testing.T) {
	cases := map[string]struct {
		param    []int
		expected [2]int
	}{
		"simplest case":  {[]int{1, 2}, [2]int{2, 1}},
		"already sorted": {[]int{1, 2, 3, 5}, [2]int{5, 3}},
		"reversed":       {[]int{5, 4, 3, 2, 1}, [2]int{5, 4}},
	}
	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			actual := misc.TwoOldestAges(testCase.param)
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
