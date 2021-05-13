/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	cur.Next = head
	for cur.Next != nil && cur.Next.Next != nil {
		// swap Next.Next & Next
		sec := cur.Next.Next
		cur.Next.Next = sec.Next
		sec.Next = cur.Next
		cur.Next = sec

		// move to next.next
		cur = cur.Next.Next
	}
	return dummy.Next
}

// @lc code=end

