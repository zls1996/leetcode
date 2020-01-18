/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(s string) int {
	if len(s) == 0 {
		return 0
	}

	var pow float64
	var result float64
	for index := len(s) -1; index >= 0; index-- {
		val := float64(s[index]-'A'+1)
		result += val * math.Pow(26, pow)
		pow += 1
	}

	return int(result)
    
}
// @lc code=end

