package misc

// https://leetcode.com/problems/boats-to-save-people/

import (
	"testing"
)

func TestNumRescueBoats(t *testing.T) {
	cases := map[string]struct {
		people   []int
		limit    int
		expected int
	}{
		"firstCase": {
			[]int{1, 2},
			3,
			1,
		},
		"secondCase": {
			[]int{3, 2, 2, 1},
			3,
			3,
		},
		"thirdCase": {
			[]int{3, 5, 3, 4},
			5,
			4,
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := NumRescueBoats(tt.people, tt.limit)
			expected := tt.expected
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
