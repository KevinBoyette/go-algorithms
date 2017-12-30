package str_test

// Named str because strings and string are keywords
import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/str"
)

func TestSimpleCompression(t *testing.T) {
	testTable := []struct {
		testName  string
		testParam string
		expected  string
	}{
		{"testing SimpleCompression('aaaabbbccdaa')", "aaaabbbccdaa", "a4b3c2d1a2"},
		{"testing SimpleCompression()", "", ""},
		{"testing SimpleCompression('a')", "a", "a1"},
		{"testing SimpleCompression('AAAABBBCCDAA')", "AAAABBBCCDAA", "a4b3c2d1a2"},
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			actual := str.SimpleCompression(test.testParam)
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
