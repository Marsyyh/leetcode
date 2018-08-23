package quickSort

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestQuickSort(t *testing.T) {
	input := []int{5, 3, 2, 8, 4, 7, 1, 6}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
	testHelper.AssertSliceEqual(expect, quickSort(input), t)
}
