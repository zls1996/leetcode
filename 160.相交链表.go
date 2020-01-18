/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB

	if a == nil || b == nil {
		return nil
	}


	for a != nil && b != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}

	if a == nil {
		a = headB
		for a != nil && b != nil {
			if a == b {
				return a
			}
			a = a.Next
			b = b.Next
			if b == nil {
				b = headA
			}
		}
	}

	

	if b == nil {
		b = headA
		for a != nil && b != nil {
			if a == b {
				return a
			}
			a = a.Next
			if a == nil {
				a = headB
			}
			b = b.Next
		}
	}

	

	return nil

}
// @lc code=end

