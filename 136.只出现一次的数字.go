/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	var result int
	
	for index := 0; index < len(nums); index++ {
		result = result ^ nums[index]
	}

	return result
}
// @lc code=end

