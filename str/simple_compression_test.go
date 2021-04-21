package str_test

// Named str because strings and string are keywords
import (
	"testing"

	"kevinboyette/algorithms/str"
)

func TestSimpleCompression(t *testing.T) {
	cases := map[string]struct {
		testParam string
		expected  string
	}{
		"multiple instances of letters":             {"aaaabbbccdaa", "a4b3c2d1a2"},
		"empty string":                              {"", ""},
		"single letter":                             {"a", "a1"},
		"multiple instances of letters in all caps": {"AAAABBBCCDAA", "a4b3c2d1a2"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := str.SimpleCompression(tc.testParam)
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
