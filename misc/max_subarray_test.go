package misc_test

import (
	subArray "kevinboyette/algorithms/misc"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := map[string]struct {
		param    []int
		expected int
	}{
		"first case": {
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		"second case": {
			[]int{0},
			0,
		},
		"third case": {
			[]int{-1},
			-1,
		},
		"fourth case": {
			[]int{-100000},
			-100000,
		},
	}
	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			actual := subArray.MaxSubArray(testCase.param)
			expected := testCase.expected
			if actual != expected {
				t.Errorf("During %s; expected %v and got %v",
					name,
					expected,
					actual,
				)
			}

		})
	}
}
