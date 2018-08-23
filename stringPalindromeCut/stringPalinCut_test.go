package stringPalindromeCut

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestMinCuts(t *testing.T) {
	re := MinCuts("ababbbabbababa")
	re1 := MinCuts("abba")
	re2 := MinCuts("addbda")
	testHelper.AssertEqual(re, 3, t)
	testHelper.AssertEqual(re1, 0, t)
	testHelper.AssertEqual(re2, 3, t)
}
