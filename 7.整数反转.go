/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	const INT_MAX = math.MaxInt32
	const INT_MIN = math.MinInt32

	var result int = 0

	for index := x % 10; x != 0; index = x % 10 {
		result = result * 10 + index
		if result > INT_MAX || result < INT_MIN {
			return 0
		}
		x = (x-index) / 10
	}


	return result

}
// @lc code=end

