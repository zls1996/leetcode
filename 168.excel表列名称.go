/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(n int) string {
	arr := make([]byte , 0)
	
	for n != 0 {
		val := n % 26
		if val == 0 {
			val = 26
			n -= 26
		} else {
			n -= val
		}
		arr = append(arr, byte(val-1 + int('A')))
		n = n / 26
	}

	begin := 0
	end := len(arr) - 1

	for end > begin {
		temp := arr[begin]
		arr[begin] = arr[end]
		arr[end] = temp
		end--
		begin++
	}

	return string(arr[:])


}
// @lc code=end

