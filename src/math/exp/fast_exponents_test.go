package exp_test

import (
	"testing"

	"algorithms/src/math/exp"
)

func TestFastExponents(t *testing.T) {
	cases := map[string]struct {
		testParamX int
		testParamY int
		expected   int
	}{
		"2^0": {2, 0, 1},
		// TODO {"testing FastExponents(2,0)", 2, -1, .5},
		// TODO {"testing FastExponents(2,0)", 2, -2, .25},
		"2^3":    {2, 3, 8},
		"3^3":    {3, 3, 27},
		"16^8":   {16, 8, 4294967296},
		"1000^1": {1000, 1, 1000},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := exp.FastExponents(tc.testParamX, tc.testParamY)
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
