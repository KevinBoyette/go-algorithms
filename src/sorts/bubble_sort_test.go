package sorts_test

import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/sorts"
)

func TestBubbleSort(t *testing.T) {
	cases := map[string]struct {
		testParam []int
		expected  []int
	}{
		"BubbleSort([]int{})":          {[]int{}, []int{}},
		"BubbleSort([]int{1})":         {[]int{1}, []int{1}},
		"BubbleSort([]int{5,4,3,2,1})": {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"BubbleSort([]int{2,1,3,5,4})": {[]int{2, 1, 3, 5, 4}, []int{1, 2, 3, 4, 5}},
		"BubbleSort([]int{1,2,3,4,5})": {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := sorts.BubbleSort(tc.testParam)
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
