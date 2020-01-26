/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	for node.Next != nil {
		next := node.Next
		val := node.Val
		node.Val = next.Val
		next.Val = val
		if next.Next == nil {
			break
		}
		node = node.Next
	}

	node.Next = nil
	

}
// @lc code=end

