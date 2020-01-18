/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	resultArray := make([][] int , maxDepth(root))
	orderByLevel(0, root, resultArray)

	begin := 0 
	end := len(resultArray) - 1

	for begin < end {
		temp := resultArray[begin]
		resultArray[begin] = resultArray[end]
		resultArray[end] = temp
		begin++
		end--
	}

	return resultArray
}

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

func orderByLevel(level int, root *TreeNode, resultArray [][] int) {
	if nil == root {
		return
	}
	if level == len(resultArray) {
		resultArray = append(resultArray, make([]int, 1))
	}
	array := resultArray[level]
	array = append(array, root.Val)
	resultArray[level] = array
	level++
	orderByLevel(level, root.Left, resultArray)
	orderByLevel(level, root.Right, resultArray)
	
}
// @lc code=end

