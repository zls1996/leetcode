/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {

	length := 32
	result := 0
	for length > 0 {
		val := (int)(num & 1)
		result += val
		num = num >> 1
		length--
	}

	return result
    
}
// @lc code=end

