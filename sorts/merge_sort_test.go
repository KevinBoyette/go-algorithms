package sorts

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := map[string]struct {
		given []int
		want  []int
	}{
		"empty":                                 {[]int{}, []int{}},
		"already sorted":                        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"single element":                        {[]int{1}, []int{1}},
		"negative elements":                     {[]int{-1, -2, 2, 0, 3}, []int{-2, -1, 0, 2, 3}},
		"reversed slice":                        {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"multiple versions of similar elements": {[]int{5, 4, 3, 4, 1, 2, 1, 3, 4, 5}, []int{1, 1, 2, 3, 3, 4, 4, 4, 5, 5}},
		"all the same element":                  {[]int{1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1}},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := MergeSort(tt.given)
			if len(tt.given) != len(actual) {
				t.Errorf("During %s; expected len(%v) and got len(%v)",
					name,
					len(tt.want),
					len(actual),
				)
			}
			for i, v := range actual {
				if v != tt.want[i] {
					t.Errorf("During %s; expected %v and got %v",
						name,
						tt.want,
						actual,
					)
				}
			}
		})
	}
}

func TestMerge(t *testing.T) {
	cases := map[string]struct {
		givenOne []int
		givenTwo []int
		want     []int
	}{
		"empty slice":                    {[]int{}, []int{}, []int{}},
		"compare slices already sorted":  {[]int{1, 2}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		"one element in the left slice":  {[]int{1}, []int{}, []int{1}},
		"one element in the right slice": {[]int{}, []int{1}, []int{1}},
		"right side has smaller values":  {[]int{2, 0, 3}, []int{-1, -2}, []int{-1, -2, 2, 0, 3}},
		"left side has smaller values":   {[]int{-1, -2}, []int{2, 0, 3}, []int{-1, -2, 2, 0, 3}},
		"normal looking slices":          {[]int{2, 0, 3}, []int{-1, 0, -2}, []int{-1, 0, -2, 2, 0, 3}},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := Merge(tt.givenOne, tt.givenTwo)
			if len(tt.givenOne)+len(tt.givenTwo) != len(actual) {
				t.Errorf("During %s; expected len(%v) and got len(%v)",
					name,
					len(tt.want),
					len(actual),
				)
			}
			for i, v := range actual {
				if v != tt.want[i] {
					t.Errorf("During %s; expected %v and got %v",
						name,
						tt.want,
						actual,
					)
				}
			}
		})
	}
}
