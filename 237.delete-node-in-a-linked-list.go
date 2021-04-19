/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    if node == nil {
        return
    }
    n := node.Next
    // node is not tail
    if n != nil {
        node.Val = n.Val
        node.Next = n.Next
        n.Next = nil
    }    
    // node is tail
}

