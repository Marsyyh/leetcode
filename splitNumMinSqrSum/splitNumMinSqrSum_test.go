package splitNumMinSqrSum

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestMinSqrSum(t *testing.T) {
	testHelper.AssertEqual(MinSqrSum(4), 1, t)
	testHelper.AssertEqual(MinSqrSum(10), 2, t)
	testHelper.AssertEqual(MinSqrSum(22), 3, t)
}
