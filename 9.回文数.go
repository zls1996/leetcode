/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}
	
	var tempNum int = x
	y := int(0)

	for tempNum != 0 {
		t := tempNum % 10
		y = y * 10 + t
		tempNum = (tempNum - t) / 10
	}
	return x == y

}
// @lc code=end

