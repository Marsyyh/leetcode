package firstMissPositive

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestFirstMissPositive(t *testing.T) {
	v := []int{3, 4, -1, 1}
	r := firstMissingPositive(v)
	testHelper.AssertEqual(2, r, t)
}

func TestFirstMissPositive2(t *testing.T) {
	v := []int{3, 4, -1, 1}
	r := firstMissingPositive2(v)
	testHelper.AssertEqual(2, r, t)
}
