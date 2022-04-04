package euler

import (
	"testing"
)

func TestEulerTotient(t *testing.T) {
	cases := map[string]struct {
		testParam uint
		expected  uint
	}{
		"Totient(10)": {10, 4},
		"Totient(9)":  {9, 6},
		"Totient(5)":  {5, 4},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := Totient(tc.testParam)
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
