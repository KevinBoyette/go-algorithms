package exp

import "testing"

func TestFastExponents(t *testing.T) {
	testTable := []struct {
		testName   string
		testParamX int
		testParamY int
		expected   int
	}{
		{"testing FastExponents(2,0)", 2, 0, 1},
		// TODO {"testing FastExponents(2,0)", 2, -1, .5},
		// TODO {"testing FastExponents(2,0)", 2, -2, .25},
		{"testing FastExponents(2,3)", 2, 3, 8},
		{"testing FastExponents(3,3)", 3, 3, 27},
		{"testing FastExponents(16,8)", 16, 8, 4294967296},
		{"testing FastExponents(1000,1)", 1000, 1, 1000},
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			actual := FastExponents(test.testParamX, test.testParamY)
			expected := test.expected
			if actual != expected {
				t.Errorf("During %s; expected %v and got %v",
					test.testName,
					test.expected,
					actual,
				)
			}
		})
	}

}
