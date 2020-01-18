/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	index := 0
	currentValue := nums[index]
	
	for i := 1; i < len(nums); i++ {
		if nums[i] > currentValue {
			currentValue = nums[i]
			index++
			nums[index] = currentValue
		}
	}

	nums = nums[: index+1]
	return len(nums)
}
// @lc code=end

