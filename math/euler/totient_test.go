package euler

import (
	"testing"
)

func TestEulerTotient(t *testing.T) {
	cases := map[string]struct {
		given uint
		want  uint
	}{
		"Totient(10)": {10, 4},
		"Totient(9)":  {9, 6},
		"Totient(5)":  {5, 4},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := Totient(tt.given)
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
