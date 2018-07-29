package twoSum

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestTwoSum(t *testing.T) {
	test := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(test, target)
	expect := []int{0, 1}
	testHelper.AssertSliceEqual(expect, res, t)
}
