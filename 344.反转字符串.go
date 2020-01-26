/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte)  {
	beginIndex := 0 
	endIndex := len(s) - 1

	for beginIndex < endIndex {
		temp := s[beginIndex]
		s[beginIndex] = s[endIndex]
		s[endIndex] = temp
		beginIndex++
		endIndex--
	}
}
// @lc code=end

