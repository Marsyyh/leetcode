/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	// init
	var stack []*TreeNode
	stack = append(stack, root)

	//gen

}

func pop(stack []*TreeNode) (*TreeNode, bool) {
	if len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]
		return node, true
	}
	return nil, false
}

// @lc code=end

