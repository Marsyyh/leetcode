/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
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
func longestUnivaluePath(root *TreeNode) int {
	res := 0
	_ = Triverse(root, &res)
	return res
}

func Triverse(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := Triverse(root.Left, res)
	r := Triverse(root.Right, res)
	if l == 0 && r == 0{
		return 1
	}
	if l == 0 && root.Val == root.Right.Val {
		return r + 1
	} else {
		if *res < r {
			*res = r
		}
		return 1
	}
	if r == 0 && root.Val == root.Left.Val {
		return l + 1
	} else {
		if *res < l {
			*res = l
		}
		return 1
	}
	if root.Val == root.Left.Val && root.Val == root.Right.Val {
		cur := l + r + 1
		if *res < cur {
			*res = cur
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}

	
	
}
// @lc code=end

