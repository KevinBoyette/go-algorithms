package fn_test

import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/fn"
)

func TestMapFunction(t *testing.T) {
	cases := mapTestTable()
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := fn.Map(tc.testFunc, tc.testSlice)
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
