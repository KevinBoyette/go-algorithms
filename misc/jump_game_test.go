package misc

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	cases := map[string]struct {
		input  []int
		output bool
	}{
		"firstCase": {
			[]int{2, 3, 1, 1, 4},
			true,
		},
		"secondCase": {
			[]int{3, 2, 1, 0, 4},
			false,
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := CanJump(tt.input)
			expected := tt.output
			if actual != expected {
				t.Errorf("During %s; expected %v and got %v",
					name,
					expected,
					actual,
				)
			}
		})
	}
}
