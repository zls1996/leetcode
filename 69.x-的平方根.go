/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	return calculate(1, x ,x)
}

func calculate(begin int, end int, target int) int {
	if begin == end-1 {
		return begin
	}
	mid := (begin+end) / 2
	value := mid*mid
	if value > target {
		return calculate(begin, mid , target)
	}

	if value < target {
		return calculate(mid, end, target)
	}

	return mid

}
// @lc code=end

