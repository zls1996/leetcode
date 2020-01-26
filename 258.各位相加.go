/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
    return ((num - 1) % 9) + 1
}
// @lc code=end

