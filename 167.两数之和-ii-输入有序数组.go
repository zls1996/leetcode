/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
    if len(numbers) == 0 {
		return make([]int, 0)
	}

	begin := 0
	end := len(numbers) -1 

	for begin < end {
		val := numbers[begin] + numbers[end]
		if val > target {
			end--
			continue
		}
		if val < target {
			begin++
			continue
		}

		return []int {begin+1, end+1}
	}

	return make([]int, 0)
}
// @lc code=end

