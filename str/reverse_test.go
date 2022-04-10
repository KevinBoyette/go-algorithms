package str

// Named str because strings and string are keywords
import (
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := map[string]struct {
		given string
		want  string
	}{
		"multiple words": {"hello world!", "!dlrow olleh"},
		"palindrome":     {"racecar", "racecar"},
		"empty string":   {"", ""},
		"single letter":  {"a", "a"},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := Reverse(tt.given)
			if actual != tt.want {
				t.Errorf("During %s; expected %v and got %v",
					name,
					tt.want,
					actual,
				)
			}
		})
	}
}
