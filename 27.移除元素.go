/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}

	begin := 0
	end := len(nums) - 1

	for end > begin {
		if (nums[end] == val) {
			end--
			continue
		}
		if (nums[begin] != val) {
			begin++
			continue
		}
		temp := nums[end]
		nums[end] = nums[begin]
		nums[begin] = temp
	}

	if nums[end] == val {
		return end
	}
	return end+1

}
// @lc code=end

