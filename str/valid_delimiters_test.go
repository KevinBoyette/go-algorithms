package str_test

import (
	"testing"

	"kevinboyette/algorithms/str"
)

func TestValidatingDelimiters(t *testing.T) {
	cases := map[string]struct {
		testParam string
		expected  bool
	}{
		"a pair of parenthesis":      {"()", true},
		"matching, but nested":       {"", true},
		"matching, but not nested":   {"()[]{}", true},
		"do not match":               {"[[", false},
		"not matching and different": {"(]", false},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := str.ValidDelimiters(tc.testParam)
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
