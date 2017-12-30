package sorts_test

import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/sorts"
)

func TestMergeSort(t *testing.T) {
	cases := map[string]struct {
		testParam []int
		expected  []int
	}{
		"MergeSort([]int{})":                    {[]int{}, []int{}},
		"MergeSort([]int{1,2,3,4,5})":           {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"MergeSort([]int{1})":                   {[]int{1}, []int{1}},
		"MergeSort([]int{-1, -2, 2, 0, 3})":     {[]int{-1, -2, 2, 0, 3}, []int{-2, -1, 0, 2, 3}},
		"MergeSort([]int{5,4,3,2,1})":           {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"MergeSort([]int{5,4,3,4,1,2,1,3,4,5})": {[]int{5, 4, 3, 4, 1, 2, 1, 3, 4, 5}, []int{1, 1, 2, 3, 3, 4, 4, 4, 5, 5}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := sorts.MergeSort(tc.testParam)
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

func TestMerge(t *testing.T) {
	cases := map[string]struct {
		testParamX []int
		testParamY []int
		expected   []int
	}{
		"Merge([]int{}, []int{})":            {[]int{}, []int{}, []int{}},
		"Merge([]int{1,2}, []int{3,4,5})":    {[]int{1, 2}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"Merge([]int{1}, []int{})":           {[]int{1}, []int{}, []int{1}},
		"Merge([]int{2,0,3},[]int{-1, -2})":  {[]int{2, 0, 3}, []int{-1, -2}, []int{-1, -2, 2, 0, 3}},
		"Merge([]int{2,0,3},[]int{-1,0 -2})": {[]int{2, 0, 3}, []int{-1, 0, -2}, []int{-1, 0, -2, 2, 0, 3}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := sorts.Merge(tc.testParamX, tc.testParamY)
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
