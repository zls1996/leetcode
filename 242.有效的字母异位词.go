/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	sArr := make([]int, 26)
	tArr := make([]int, 26)

	for index := 0; index < len(s); index++ {
		val := s[index]
		sArr[int(val - 'a')] += 1
	}

	for index := 0; index < len(t); index++ {
		val := t[index]
		tArr[int(val - 'a')] += 1
	}

	for index := 0; index < len(sArr); index++ {
		if sArr[index] != tArr[index] {
			return false
		}
	}

	return true
}
//len(sArr)code=end
