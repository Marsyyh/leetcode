package findClosestElements

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestFindClosestElements(t *testing.T) {
	test := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	k := 3
	x := 5
	expect := []int{3, 3, 4}
	result := findClosestElements(test, k, x)
	testHelper.AssertSliceEqual(expect, result, t)
}
