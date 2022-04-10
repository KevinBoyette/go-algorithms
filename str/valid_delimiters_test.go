package str

import (
	"testing"
)

func TestValidatingDelimiters(t *testing.T) {
	cases := map[string]struct {
		given string
		want  bool
	}{
		"a pair of parenthesis":      {"()", true},
		"matching, but nested":       {"", true},
		"matching, but not nested":   {"()[]{}", true},
		"do not match":               {"[[", false},
		"not matching and different": {"(]", false},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := ValidDelimiters(tt.given)
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
