/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 分析：这道题使用recursion的思路可以分为3种情况
 1. 左树或右树有child，返回child，因为对自己来说，自己就是最小parent
 2. 左右都有，返回自己，因为左右都有的情况，自己是他们的最小祖先
 3. 左右都没有，返回nil，因为没有祖先
 
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// recursion end & edge condition
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	}
	return right
}