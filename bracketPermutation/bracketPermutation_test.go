package bracketPermutation

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestPermu(t *testing.T) {
	expect := []string{"((()))", "(())()", "()()()", "()(())", "(()())"}
	testHelper.AssertSliceStringEqual(Permu(3), expect, t)
}
