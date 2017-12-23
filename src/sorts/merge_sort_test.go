package sorts

import "testing"

func TestMergeSort(t *testing.T) {
	testTable := []struct {
		testName  string
		testParam []int
		expected  []int
	}{
		{"testing MergeSort([]int{})", []int{}, []int{}},
		{"testing MergeSort([]int{1,2,3,4,5})", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"testing MergeSort([]int{1})", []int{1}, []int{1}},
		{"testing MergeSort([]int{5,4,3,2,1})", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"testing MergeSort([]int{5,4,3,4,1,2,1,3,4,5})", []int{5, 4, 3, 4, 1, 2, 1, 3, 4, 5}, []int{1, 1, 2, 3, 3, 4, 4, 4, 5, 5}},
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			actual := MergeSort(test.testParam)
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
