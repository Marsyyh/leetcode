/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

	k list, list long n
	heap: traverse all nk elements O(nk), k items in heap O(lgk)
	Time: O(nklgk), Space: heap O(k), total O(k)
	twoMerge, every layer traverse O(nk), lgk level O(lgk)
	Time: O(nklgk), Space: every layer O(nk), O(lgk) layer, total O(nklgk)
*/
func mergeKLists(lists []*ListNode) *ListNode {
	// corner cases
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	var merge *ListNode
	for len(lists) > 1 {
		merge = twoMerge(lists[0], lists[1])
		lists = append(lists, merge)[2:]
	}
	return merge
}

func twoMerge(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}

	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}

	return dummy.Next
}

// @lc code=end

