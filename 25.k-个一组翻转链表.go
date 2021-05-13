/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead, tail, next, _ := reverseK(head, 1, k)
	for next != nil {
		curHead, curTail, curNext, _ := reverseK(next, 1, k)
		tail.Next = curHead
		tail = curTail
		next = curNext
	}
	return newHead
}

// reverse recursion
// first node: head
// second node: tail
// third node: next iteration node
// bool: if reverse
func reverseK(head *ListNode, n, k int) (*ListNode, *ListNode, *ListNode, bool) {
	// reverse link list with k
	if n == k {
		return head, head, head.Next, true
	}

	if head.Next == nil {
		return head, head, nil, false
	}

	newHead, tail, next, ok := reverseK(head.Next, n+1, k)
	if ok {
		head.Next.Next = head
		head.Next = nil
		// if reverse return head & new tail
		return newHead, head, next, ok
	}

	// if not reverse return current head & current tail
	return head, tail, nil, false
}

// @lc code=end

