package validateBinarySearchTree

import (
	"testing"

	"github.com/Marsyyh/go-datastructure/container/tree/binarytree"
	"github.com/Marsyyh/leetcode/testHelper"
)

func TestValidateBinarySearchTree(t *testing.T) {
	test := []interface{}{4, 2, 6, 1, 3, 5, 7}
	arrNode := binarytree.NewArrayTree(test)
	root := arrNode.GetRoot()
	testHelper.AssertTrue(validateBST(root), t)
}

func TestValidateBinarySearchTreeFalse(t *testing.T) {
	test := []interface{}{4, 6, 2, 1, 3, 5, 7}
	arrNode := binarytree.NewArrayTree(test)
	root := arrNode.GetRoot()
	testHelper.AssertFalse(validateBST(root), t)
}
