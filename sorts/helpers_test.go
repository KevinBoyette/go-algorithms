package sorts

import (
	"testing"
)

type testCase struct {
	testParam []int
	expected  []int
}

type twoParamTestCase struct {
	testParamX []int
	testParamY []int
	expected   []int
}

// testTable generates all the test cases for the sorting functions
// with one parameter
func testTable() map[string]testCase {
	return map[string]testCase{
		"empty":                                 {[]int{}, []int{}},
		"already sorted":                        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"single element":                        {[]int{1}, []int{1}},
		"negative elements":                     {[]int{-1, -2, 2, 0, 3}, []int{-2, -1, 0, 2, 3}},
		"reversed slice":                        {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"multiple versions of similar elements": {[]int{5, 4, 3, 4, 1, 2, 1, 3, 4, 5}, []int{1, 1, 2, 3, 3, 4, 4, 4, 5, 5}},
		"all the same element":                  {[]int{1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1}},
	}
}

// twoParamTestTable generates all the test cases for sorting functions
// that need two parameters.
func twoParamTestTable() map[string]twoParamTestCase {
	return map[string]twoParamTestCase{
		"empty slice":                    {[]int{}, []int{}, []int{}},
		"compare slices already sorted":  {[]int{1, 2}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"one element in the left slice":  {[]int{1}, []int{}, []int{1}},
		"one element in the right slice": {[]int{}, []int{1}, []int{1}},
		"right side has smaller values":  {[]int{2, 0, 3}, []int{-1, -2}, []int{-1, -2, 2, 0, 3}},
		"left side has smaller values":   {[]int{-1, -2}, []int{2, 0, 3}, []int{-1, -2, 2, 0, 3}},
		"normal looking slices":          {[]int{2, 0, 3}, []int{-1, 0, -2}, []int{-1, 0, -2, 2, 0, 3}},
	}
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

func runTestsTwoParams(f func([]int, []int) []int, testTable map[string]twoParamTestCase, t *testing.T) {
	for name, tc := range testTable {
		t.Run(name, func(t *testing.T) {
			actual := f(tc.testParamX, tc.testParamY)
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
