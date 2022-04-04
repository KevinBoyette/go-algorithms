package sorts

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := testTable()
	runTests(QuickSortRecur, cases, t)
}
