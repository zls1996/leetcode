/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	val := 1

	for val <= n {
		if val == n {
			return true
		}
		val = val << 1
	}

	return false
}
// @lc code=end

