/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	length := len(nums)
	result := length * (length + 1) / 2
	for index := 0; index < length; index++ {
		val := nums[index]
		result -= val
	}

	return result
}
// @lc code=end

