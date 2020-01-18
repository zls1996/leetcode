/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	count := 1
	result := nums[0]
	for index := 1; index < len(nums); index++ {
		val := nums[index]
		if count == 0 {
			count++
			result = val 
			continue
		}
		if val == result {
			count++
		} else {
			count--
		}
	}

	return result
}
// @lc code=end

