package sorts

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	cases := testTable()
	runTests(BubbleSort, cases, t)
}
