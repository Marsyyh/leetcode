/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	zero := &ListNode{Val: 0}
	next := head
	cap := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = zero
		}
		if l2 == nil {
			l2 = zero
		}

		cur := l1.Val + l2.Val + cap
		cap = 0
		if cur > 9 {
			cap = 1
			cur = cur - 10
		}

		next.Next = &ListNode{Val: cur}
		next = next.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if cap != 0 {
		next.Next = &ListNode{Val: cap}
	}

	return head.Next
}

// @lc code=end

