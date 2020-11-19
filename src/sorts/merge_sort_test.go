package sorts_test

import (
	"testing"

	"algorithms/src/sorts"
)

func TestMergeSort(t *testing.T) {
	cases := testTable()
	runTests(sorts.MergeSort, cases, t)
}

func TestMerge(t *testing.T) {
	cases := twoParamTestTable()
	runTestsTwoParams(sorts.Merge, cases, t)
}
