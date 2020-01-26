/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
   if root == nil {
	   return make([] string, 0)
   } 

   result := make([] string, 0)

   result = travelsal(root, "", result)

   return result
}

func travelsal(root *TreeNode, str string, result [] string) []string {
	appendStr := fmt.Sprintf("%s%d", str, root.Val)
	if root.Left == nil && root.Right == nil {
		result = append(result, appendStr)
		return result
	}
	appendStr = appendStr + "->"
	if root.Left == nil {
		return travelsal(root.Right, appendStr, result)
	}

	if root.Right == nil {
		return travelsal(root.Left, appendStr, result)
	}

	result = travelsal(root.Left, appendStr, result)
	result = travelsal(root.Right, appendStr, result)

	return result
}
// @lc code=end

