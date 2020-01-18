/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
    return buildAsTree(0, len(nums)-1, nums)
}

func buildAsTree(begin int, end int, nums []int) *TreeNode {
	if begin > end {
		return nil
	}
	mid := (begin + end) / 2
	node := &TreeNode{
		Val: nums[mid],
		Left: buildAsTree(begin, mid-1, nums),
		Right: buildAsTree(mid+1, end, nums),
	}

	return node
}
// @lc code=end

