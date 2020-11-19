package sorts_test

import (
	"testing"

	"algorithms/src/sorts"
)

func TestQuickSort(t *testing.T) {
	cases := testTable()
	runTests(sorts.QuickSortRecur, cases, t)
}
