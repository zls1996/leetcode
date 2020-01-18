/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {

	resultMap := make(map[int] int, n)
	resultMap[1] = 1
	resultMap[2] = 2
	return getByMap(n, resultMap)
    
}

func getByMap(n int, resultMap map[int] int) int {
	value, ok := resultMap[n]
	if ok {
		return value
	}

	result := getByMap(n-1, resultMap) + getByMap(n-2, resultMap)
	resultMap[n] = result
	return result
}
// @lc code=end

