package stringSpacePermutation

import (
	"testing"

	"github.com/Marsyyh/leetcode/testHelper"
)

func TestSpacePerm(t *testing.T) {
	input := "ABC"
	expected := []string{"ABC", "AB_C", "A_BC", "A_B_C"}
	testHelper.AssertSliceStringEqual(SpacePerm(input), expected, t)
}
