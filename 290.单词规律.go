/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {
	strArray := strings.Split(str, " ")
	if len(strArray) != len(pattern) {
		return false
	}
	wordArray := make([]string, 26)
	wordMap := make(map[string] byte)
	
	for index := 0; index < len(pattern); index++ {
		byteIndex := int(pattern[index] - 'a')
		val := wordArray[byteIndex]
		if val == "" {
			wordArray[byteIndex] = strArray[index]
			mapVal, ok := wordMap[strArray[index]]
			if ok {
				if mapVal != pattern[index] {
					return false
				}
			} else {
				wordMap[strArray[index]] = pattern[index]
			}
		} else {
			if val != strArray[index] {
				return false
			}
			mapVal, ok := wordMap[val] 
			if !ok {
				return false
			}
			if pattern[index] != mapVal {
				return false
			}
		}
	}

	return true
}
// @lc code=end

