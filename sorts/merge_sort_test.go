package sorts

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := testTable()
	runTests(MergeSort, cases, t)
}

func TestMerge(t *testing.T) {
	cases := twoParamTestTable()
	runTestsTwoParams(Merge, cases, t)
}
