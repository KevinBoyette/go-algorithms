package euler

import "testing"

func TestEulerTotient(t *testing.T) {
	testTable := []struct {
		testName  string
		testParam uint
		expected  uint
	}{
		{"testing Totient(10)", 10, 4},
		{"testing Totient(9)", 9, 6},
		{"testing Totient(5)", 5, 4},
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			actual := Totient(test.testParam)
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
