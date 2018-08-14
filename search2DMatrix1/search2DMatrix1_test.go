package search2DMatrix1

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestSearch2DMatrix1(t *testing.T) {
	test := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	testHelper.AssertTrue(searchMatrix(test, 16), t)
	testHelper.AssertTrue(searchMatrix(test, 1), t)
	testHelper.AssertFalse(searchMatrix(test, 24), t)
}
