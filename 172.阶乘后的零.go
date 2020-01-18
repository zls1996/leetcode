/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {
	result := 0
	
	val := n / 5
	result += val 

	for val != 0 {
		val /= 5
		result += val
	}

	return result
}
// @lc code=end

