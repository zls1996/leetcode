/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	val := sum - root.Val

	if root.Left == nil {
		return hasPathSum(root.Right, val)
	}

	if root.Right == nil {
		return hasPathSum(root.Left, val)
	}

	return hasPathSum(root.Right, val) || hasPathSum(root.Left, val)


}
// @lc code=end

