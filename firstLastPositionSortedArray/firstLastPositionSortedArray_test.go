package firstLastPositionSortedArray

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestFirstLastPositionSortedArray(t *testing.T) {
	test := []int{1, 2, 3, 3, 3, 3}
	expect := []int{2, 5}
	target := 3
	testHelper.AssertSliceEqual(expect, searchRange(test, target), t)
	test = []int{1}
	expect = []int{-1, -1}
	target = 3
	testHelper.AssertSliceEqual(expect, searchRange(test, target), t)
}

func TestFirstLastPositionSortedArray1(t *testing.T) {
	test := []int{1, 2, 3, 3, 3, 3}
	expect := []int{2, 5}
	target := 3
	testHelper.AssertSliceEqual(expect, searchRange1(test, target), t)
	test = []int{1}
	expect = []int{-1, -1}
	target = 3
	testHelper.AssertSliceEqual(expect, searchRange1(test, target), t)
}
