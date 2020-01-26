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
	if head == nil || head.Next == nil{
		return true
	}

	slow := head
	fast := head
	var reverseNode *ListNode 

	for fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast.Next == nil {
			break
		}
	}


	reverseNode = slow.Next
	slow.Next = nil
	finalNode := reverseNode
	for reverseNode.Next != nil {
		temp := reverseNode.Next
		reverseNode.Next = temp.Next
		temp.Next = finalNode
		finalNode = temp
	}

	for head != nil && finalNode != nil {
		if head.Val != finalNode.Val {
			return false
		}
		head = head.Next
		finalNode = finalNode.Next
	}

	return true

	
}

// @lc code=end

