package misc

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := map[string]struct {
		given []int
		want  int
	}{
		"first case": {
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		"second case": {
			[]int{0},
			0,
		},
		"third case": {
			[]int{-1},
			-1,
		},
		"fourth case": {
			[]int{-100000},
			-100000,
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := MaxSubArray(tt.given)
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
