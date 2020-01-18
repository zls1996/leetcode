/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i <= (len(haystack)-len(needle)); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			if (haystack[i+j] != needle[j] ) {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
    
}
// @lc code=end

