package misc_test

// https://leetcode.com/problems/boats-to-save-people/

import (
	boats "kevinboyette/algorithms/misc"
	"testing"
)

// TODO: move these tests into test table
func TestNumRescueBoatsFirstCase(t *testing.T) {
	people := []int{1, 2}
	limit := 3
	actual := boats.NumRescueBoats(people, limit)
	expected := 1
	if actual != expected {
		t.Errorf("actual %d != expected %d", actual, expected)
	}
}
func TestNumRescueBoatsSecondCase(t *testing.T) {
	people := []int{3, 2, 2, 1}
	limit := 3
	actual := boats.NumRescueBoats(people, limit)
	expected := 3
	if actual != expected {
		t.Errorf("actual %d != expected %d", actual, expected)
	}
}
func TestNumRescueBoatsThirdCase(t *testing.T) {
	people := []int{3, 5, 3, 4}
	limit := 5
	actual := boats.NumRescueBoats(people, limit)
	expected := 4
	if actual != expected {
		t.Errorf("actual %d != expected %d", actual, expected)
	}
}
