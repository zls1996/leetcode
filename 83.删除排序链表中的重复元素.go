/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first := head
	if first.Next == nil {
		return head
	}
	second := first.Next
	for second != nil {
		if second.Val == first.Val {
			second = second.Next
			first.Next = second
		} else {
			first = second
			second = second.Next
		}
	}

	return head
}
// @lc code=end

