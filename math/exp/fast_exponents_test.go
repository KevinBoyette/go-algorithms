package exp

import (
	"testing"
)

func TestFastExponents(t *testing.T) {
	cases := map[string]struct {
		given [2]int
		want  int
	}{
		"2^0": {[2]int{2, 0}, 1},
		// TODO {"testing FastExponents(2,0)", 2, -1, .5},
		// TODO {"testing FastExponents(2,0)", 2, -2, .25},
		"2^3":    {[2]int{2, 3}, 8},
		"3^3":    {[2]int{3, 3}, 27},
		"16^8":   {[2]int{16, 8}, 4294967296},
		"1000^1": {[2]int{1000, 1}, 1000},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := FastExponents(tt.given[0], tt.given[1])
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
