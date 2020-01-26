/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result := getLatestParent(p, q)
	if result != nil {
		return p
	}
	result = getLatestParent(q, p)
	if result != nil {
		return q
	}
	parentP := getLatestParent(root, p)
	parentQ := getLatestParent(root, q)
	if parentP == nil || parentQ == nil {
		return nil
	}
	if parentP == parentQ {
		return parentP
	}
	return lowestCommonAncestor(root, parentP, parentQ)
}

func getLatestParent(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node == root {
		return root
	}
	parent := root

	if parent.Left == node || parent.Right == node {
		return parent
	}

	left := getLatestParent(parent.Left, node)
	if left != nil {
		return left
	}

	right := getLatestParent(parent.Right, node)
	if right != nil {
		return right
	}

	return nil
}
// @lc code=end

