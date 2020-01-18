/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	first := head
	second := first.Next

	first.Next = nil

	for second != nil {
		temp := second
		second = second.Next
		temp.Next = first
		first = temp
	}

	return first

}
// @lc code=end

