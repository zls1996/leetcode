/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
    if len(nums) == 0 {
		return 
	}
	nonZeroIndex := 0
	for index := 0; index < len(nums); index++ {
		val := nums[index]
		if val != 0 {
			nums[nonZeroIndex] = val
			nonZeroIndex++
		}
	}

	for nonZeroIndex < len(nums) {
		nums[nonZeroIndex] = 0
		nonZeroIndex++
	}

}
// @lc code=end

