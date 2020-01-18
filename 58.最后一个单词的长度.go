/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	begin := -1
	end := -1

	for index := len(s) - 1; index >= 0; index-- {
		if s[index] == ' ' {
			if end != -1 {
				begin = index
				break
			}
		} else {
			if end == -1 {
				end = index
			}
		}
	}

	if end == -1 {
		return 0
	}

	if begin == -1 {
		return end + 1
	}

	return end - begin

}
// @lc code=end

