/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	val := head.Val
	head = head.Next
	if head == nil {
		return true
	}

	for head != nil {
		val = val ^ head.Val
		head = head.Next
	}

	if val != 0 {
		return false
	}

	return true

	
}

// @lc code=end

