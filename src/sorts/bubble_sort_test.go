package sorts_test

import (
	"testing"

	"github.com/KevinBoyette/GoAlgorithms/src/sorts"
)

func TestBubbleSort(t *testing.T) {
	cases := testTable()
	runTests(sorts.BubbleSort, cases, t)
}
