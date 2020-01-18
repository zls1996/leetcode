/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	return doPlus(digits, len(digits)-1)
}

func doPlus(digits []int, index int) []int {
	if index < 0 {
		return digits
	}

	if digits[index] < 9 {
		digits[index] += 1
		return digits
	}

	digits[index] = 0

	if index == 0 {
		result := make([]int, len(digits) + 1)
		result[0] = 1
		for i := 1; i < len(result); i++ {
			result[i] = digits[i-1]
		}
		return result
	}

	index--
	return doPlus(digits, index)
}
// @lc code=end

