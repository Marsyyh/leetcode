package treeCousins

import (
	"testing"

	"github.com/Marsyyh/go-datastructure/container/tree/binarytree"
	"github.com/Marsyyh/leetcode/testHelper"
)

func TestIsCousin(t *testing.T) {
	test := []interface{}{6, 3, 5, 7, 8, 1, 2}
	arrNode := binarytree.NewArrayTree(test)
	root := arrNode.GetRoot()
	testHelper.AssertTrue(IsCousin(root, 7, 1), t)
	testHelper.AssertFalse(IsCousin(root, 3, 5), t)
	testHelper.AssertFalse(IsCousin(root, 7, 5), t)
}
