/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	recursion(dummy, n)
	return dummy.Next
}

func recursion(head *ListNode, n int) int {
	// end
	if head.Next == nil {
		return 1
	}

	// get current
	cur := recursion(head.Next, n)
	if cur == n {
		tmp := head.Next
		head.Next = head.Next.Next
		tmp.Next = nil
	}
	return cur + 1
}

// @lc code=end

