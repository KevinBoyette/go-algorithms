package fn_test

import (
	"algorithms/src/fn"
	"testing"
)

func TestReduceFunction(t *testing.T) {
	cases := reduceTestTable()
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := fn.Reduce(tc.testFunc, tc.testSlice)
			expected := tc.expected
			if actual != expected {
				t.Errorf("During %s; expected %v and got %v",
					name,
					tc.expected,
					actual,
				)
			}

		})
	}
}
