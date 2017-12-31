package sorts_test

import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/sorts"
)

func TestQuickSort(t *testing.T) {
	cases := testTable()
	runTests(sorts.QuickSortRecur, cases, t)
}
