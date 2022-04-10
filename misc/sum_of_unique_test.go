package misc

import (
	"testing"
)

func TestSumOfUnique(t *testing.T) {
	cases := map[string]struct {
		given []int
		want  int
	}{
		"First Case": {
			[]int{1, 2, 3, 2},
			4,
		},
		"Second Case": {
			[]int{1, 1, 1, 1, 1},
			0,
		},
		"Third Case": {
			[]int{1, 2, 3, 4, 5},
			15,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := SumOfUniqueElements(tt.given)
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
