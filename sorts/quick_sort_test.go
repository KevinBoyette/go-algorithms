package sorts_test

import (
	"testing"

	"kevinboyette/algorithms/sorts"
)

func TestQuickSort(t *testing.T) {
	cases := testTable()
	runTests(sorts.QuickSortRecur, cases, t)
}
