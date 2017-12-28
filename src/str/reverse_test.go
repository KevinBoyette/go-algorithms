package str

// Named str because strings and string are keywords
import "testing"

func TestReverseString(t *testing.T) {
	testTable := []struct {
		testName  string
		testParam string
		expected  string
	}{
		{"testing Reverse('hello')", "hello", "olleh"},
		{"testing Reverse('racecar')", "racecar", "racecar"},
		{"testing Reverse()", "", ""},
		{"testing Reverse('a')", "a", "a"},
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			actual := Reverse(test.testParam)
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
