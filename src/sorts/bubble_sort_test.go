package sorts_test

import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/sorts"
)

func TestBubbleSort(t *testing.T) {
	testTable := []struct {
		testName  string
		testParam []int
		expected  []int
	}{
		{"testing BubbleSort([]int{})", []int{}, []int{}},
		{"testing BubbleSort([]int{1})", []int{1}, []int{1}},
		{"testing BubbleSort([]int{5,4,3,2,1})", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"testing BubbleSort([]int{2,1,3,5,4})", []int{2, 1, 3, 5, 4}, []int{1, 2, 3, 4, 5}},
		{"testing BubbleSort([]int{1,2,3,4,5})", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			actual := sorts.BubbleSort(test.testParam)
			expected := test.expected
			testCase := true
			for index := range actual {
				// This if can result in panic: index out of range
				if actual[index] != expected[index] {
					testCase = false
				}
			}
			if testCase != true {
				t.Errorf("During %s; expected %v and got %v",
					test.testName,
					test.expected,
					actual,
				)
			}
		})
	}

}
