/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	leftLength := getDepth(root.Left)
	rightLength := getDepth(root.Right)

	if leftLength == -1 || rightLength == -1 {
		return false
	}
	between := leftLength - rightLength

	if between < -1 || between > 1 {
		return false
	}

	return true
	
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLength := getDepth(root.Left)
	rightLength := getDepth(root.Right)

	if leftLength == -1 || rightLength == -1 {
		return -1
	}

	between := leftLength - rightLength

	if between < -1 || between > 1 {
		return -1
	}

	if leftLength > rightLength {
		return leftLength + 1
	}

	return rightLength + 1
}


// @lc code=end

