package stringsArrCircular

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestIsCircular(t *testing.T) {
	input := []string{"aaa", "bba", "aca", "aab"}
	testHelper.AssertTrue(IsCircular(input), t)
	input = []string{"aaa", "bbb", "ccc"}
	testHelper.AssertFalse(IsCircular(input), t)
}
