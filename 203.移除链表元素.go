/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	first := head

	for first != nil && first.Val == val {
		first = first.Next
	}

	if first == nil {
		return first
	}

	result := first
	second := first.Next

	for second != nil {
		if second.Val == val {
			first.Next = second.Next
			second = second.Next
		} else {
			second = second.Next
			first = first.Next
		}
	}
	return result
}
// @lc code=end

