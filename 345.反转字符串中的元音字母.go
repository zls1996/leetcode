/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	strArray := []byte(s)
	beginIndex := 0 
	endIndex := len(strArray) - 1
	for beginIndex < endIndex {
		for {
			val := strArray[beginIndex]
			if isVowel(val) {
				break
			}
			beginIndex++
			if beginIndex >= len(strArray) {
				break
			}
		}
	
		for {
			endVal := strArray[endIndex]
			if isVowel(endVal) {
				break
			}
			endIndex--
			if endIndex < 0 {
				break
			}
		}
		if beginIndex >= endIndex {
			break
		}
		
		temp := strArray[beginIndex]
		strArray[beginIndex] = strArray[endIndex]
		strArray[endIndex] = temp
		beginIndex++
		endIndex--
	}

	return string(strArray)


}

func isVowel(val byte) bool {
	if val == 'a' || val == 'A' {
		return true
	} 
	if val == 'e' || val == 'E' {
		return true
	}

	if val == 'i' || val == 'I' {
		return true
	}

	if val == 'o' || val == 'O' {
		return true
	}
	if val == 'u' || val == 'U' {
		return true
	}

	return false
}
// @lc code=end

