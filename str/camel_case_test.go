package str

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	cases := map[string]struct {
		given string
		want  string
	}{
		"empty string":      {"", ""},
		"simple":            {"simple", "simple"},
		"kebob-case":        {"the-weekend-warrior", "theWeekendWarrior"},
		"snake_case":        {"the_weekend_warrior", "theWeekendWarrior"},
		"caps first letter": {"The-Weekend-Warrior", "TheWeekendWarrior"},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := ToCamelCase(tt.given)

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

func BenchmarkToCamelCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ToCamelCase("the-Weekend-warrior")
	}
}
