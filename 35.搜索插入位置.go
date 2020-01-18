/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
		return 0
	}


	for index := 0; index < len(nums); index++ {
		if target <= nums[index] {
			return index
		}
	}

	return len(nums)
} 
// @lc code=end

