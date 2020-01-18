/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	time := 32
	for time > 0 {
		result = result << 1
		val := num & 1
		result += val
		
		num = num >> 1
		time--
	}
	return result
}
// @lc code=end

