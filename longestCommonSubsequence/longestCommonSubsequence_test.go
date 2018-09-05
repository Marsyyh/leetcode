package longestCommonSubsequence

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestLcs(t *testing.T) {
	testHelper.AssertEqual(Lcs("abcdefg", "acg"), 3, t)
	testHelper.AssertEqual(Lcs("aaaaa", "aa"), 2, t)
	testHelper.AssertEqual(Lcs("abcdefg", "hijk"), 0, t)
}
