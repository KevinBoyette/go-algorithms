package misc

import (
	"testing"
)

func TestSumOfUnique(t *testing.T) {
	cases := map[string]struct {
		param    []int
		expected int
	}{
		"First Case": {
			[]int{1, 2, 3, 2},
			4,
		},
		"Second Case": {
			[]int{1, 1, 1, 1, 1},
			0,
		},
		"Third Case": {
			[]int{1, 2, 3, 4, 5},
			15,
		},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			actual := SumOfUniqueElements(testCase.param)
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
