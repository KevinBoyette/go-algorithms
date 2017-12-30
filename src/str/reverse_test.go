package str_test

// Named str because strings and string are keywords
import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/str"
)

func TestReverseString(t *testing.T) {
	cases := map[string]struct {
		testParam string
		expected  string
	}{
		"Reverse('hello')":   {"hello", "olleh"},
		"Reverse('racecar')": {"racecar", "racecar"},
		"Reverse()":          {"", ""},
		"Reverse('a')":       {"a", "a"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := str.Reverse(tc.testParam)
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
