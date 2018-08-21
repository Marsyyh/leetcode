package searchInsertPosition

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestSearchInsertPosition(t *testing.T) {
	test := []int{1, 3, 5, 7}
	testHelper.AssertEqual(4, searchInsert(test, 8), t)
	testHelper.AssertEqual(3, searchInsert(test, 7), t)
	testHelper.AssertEqual(0, searchInsert(test, 1), t)
	testHelper.AssertEqual(0, searchInsert(test, 0), t)
	testHelper.AssertEqual(2, searchInsert(test, 4), t)
}
