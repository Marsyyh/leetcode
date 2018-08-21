package shuffleArray

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestShuffleArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	obj := Constructor(nums)
	shaffle := obj.Shuffle()
	reset := obj.Reset()
	expect := []int{2, 5, 1, 3, 4}
	testHelper.AssertSliceEqual(expect, shaffle, t)
	testHelper.AssertSliceEqual(nums, reset, t)
}
