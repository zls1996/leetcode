/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	leftLength := maxDepth(root.Left)
	rightLength := maxDepth(root.Right)

	if leftLength >= rightLength {
		return 1 + leftLength
	}

	return 1 + rightLength

}
// @lc code=end

