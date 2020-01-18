/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var result *ListNode

	if l1.Val > l2.Val {
		result = l2
		l2 = l2.Next
	} else {
		result = l1
		l1 = l1.Next
	}

	var final = result

	result.Next = nil

	for (l1 != nil && l2 != nil) {
		var temp *ListNode
		if l1.Val > l2.Val {
			temp = l2
			l2 = l2.Next
			
		} else {
			temp = l1
			l1 = l1.Next
		}

		result.Next = temp
		result = result.Next
	}

	if l2 != nil {
		result.Next = l2 
	}

	if l1 != nil {
		result.Next = l1
	}

	return final
}
// @lc code=end

