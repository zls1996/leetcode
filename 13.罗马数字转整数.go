/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	romanMap := make(map[byte] int)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000

	if len(s) == 1 {
		return romanMap[s[0]]
	}

	var result int = 0

	var index int = 0

	for index < len(s)-1 {
		first := romanMap[s[index]]
		second := romanMap[s[index+1]]
		if first >= second {
			result += first
		} else {
			result += second - first
			index++
		}
		index++
	}

	if index != len(s) { 
		result = result + romanMap[s[len(s)-1]]
	}

	return result
}
// @lc code=end

