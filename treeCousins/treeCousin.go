package treeCousins

import . "github.com/Marsyyh/go-datastructure/container/tree/binarytree"

func IsCousin(node *TreeNode, a, b int) bool {
	re := false
	recur(node, a, b, 0, &re)
	return re
}

func recur(node *TreeNode, a, b, level int, re *bool) int {
	if node == nil {
		return -1
	}
	if node.Val == a || node.Val == b {
		return level
	}

	left := recur(node.Left, a, b, level+1, re)
	right := recur(node.Right, a, b, level+1, re)
	if left == right && left != -1 &&
		(left-1 != level ||
			right-1 != level) {
		*re = true
	}
	if left != -1 {
		return left
	}
	return right
}
