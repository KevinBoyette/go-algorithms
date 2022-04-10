package str

// Named str because strings and string are keywords
import (
	"testing"
)

func TestSimpleCompression(t *testing.T) {
	cases := map[string]struct {
		given string
		want  string
	}{
		"multiple instances of letters":             {"aaaabbbccdaa", "a4b3c2d1a2"},
		"empty string":                              {"", ""},
		"single letter":                             {"a", "a1"},
		"multiple instances of letters in all caps": {"AAAABBBCCDAA", "a4b3c2d1a2"},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := SimpleCompression(tt.given)
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
