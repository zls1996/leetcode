/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0
	temp := 0
	for index := 0; index < len(nums); index++ {
		if index == 0 {
			max = nums[0]
			temp = nums[0]
			continue
		}
		if temp < 0 {
			if nums[index] >= 0 {
				temp = nums[index]
			} else {
				if temp < nums[index] {
					temp = nums[index]
				}
			}
		} else {
			if temp + nums[index] >= 0 {
				temp = temp + nums[index]
			} else {
				temp = 0
			}
		}

		if max < temp {
			max = temp
		}

	}

	return max
}
// @lc code=end

