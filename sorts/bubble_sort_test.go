package sorts_test

import (
	"testing"

	"kevinboyette/algorithms/sorts"
)

func TestBubbleSort(t *testing.T) {
	cases := testTable()
	runTests(sorts.BubbleSort, cases, t)
}
