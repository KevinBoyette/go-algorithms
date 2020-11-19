package sorts_test

import (
	"testing"

	"algorithms/src/sorts"
)

func TestBubbleSort(t *testing.T) {
	cases := testTable()
	runTests(sorts.BubbleSort, cases, t)
}
