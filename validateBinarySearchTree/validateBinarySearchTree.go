package validateBinarySearchTree

import (
	"fmt"

	"github.com/Marsyyh/go-datastructure/container/tree/binarytree"
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

// Time: O(n), worst case validate every node
// Space: O(height), const space for recurrsion on call stack
func validateBST(root *binarytree.TreeNode) bool {
	return isBST(root, minInt, maxInt)
}

func isBST(root *binarytree.TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	// min, cur, max
	return root.Val.(int) > min &&
		root.Val.(int) < max &&
		isBST(root.Left, min, root.Val.(int)) &&
		isBST(root.Right, root.Val.(int), max)
}

func main() {
	// [4, 2, 6, 1, 3, 5, 7]
	test := []interface{}{4, 2, 6, 1, 3, 5, 7}
	arrNode := binarytree.NewArrayTree(test)
	root := arrNode.GetRoot()
	fmt.Println(validateBST(root))
}
