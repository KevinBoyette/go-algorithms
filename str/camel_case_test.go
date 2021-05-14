package str_test

import (
	"kevinboyette/algorithms/str"
	"testing"
)

func TestToCamelCase(t *testing.T) {
	cases := map[string]struct {
		testParam string
		expected  string
	}{
		"empty string":      {"", ""},
		"simple":            {"simple", "simple"},
		"kebob-case":        {"the-weekend-warrior", "theWeekendWarrior"},
		"snake_case":        {"the_weekend_warrior", "theWeekendWarrior"},
		"caps first letter": {"The-Weekend-Warrior", "TheWeekendWarrior"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := str.ToCamelCase(tc.testParam)
			expected := tc.expected
			if actual != expected {
				t.Errorf("During %s; expected %v and got %v",
					name,
					tc.expected,
					actual,
				)
			}
		})
	}

}
