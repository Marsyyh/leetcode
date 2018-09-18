package wildcardMatching

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestWildcard(t *testing.T) {
	input := "aabbbcaa"
	// pattern1 := "aa**********"
	// pattern2 := "a?****b?c*a"
	// pattern3 := "****?****"
	// pattern4 := "**b**c**aa***"
	pattern5 := "a*b**?ca"
	// pattern6 := "**ba**c?*"
	// pattern7 := "bb**"
	// testHelper.AssertTrue(wildcard(input, pattern1), t)
	// testHelper.AssertTrue(wildcard(input, pattern2), t)
	// testHelper.AssertTrue(wildcard(input, pattern3), t)
	// testHelper.AssertTrue(wildcard(input, pattern4), t)
	testHelper.AssertFalse(wildcard(input, pattern5), t)
	// testHelper.AssertFalse(wildcard(input, pattern6), t)
	// testHelper.AssertFalse(wildcard(input, pattern7), t)
}
