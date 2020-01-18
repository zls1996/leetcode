/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	resultCache := make(map[int]int, len(nums))
	return doRub(nums, 0, resultCache)
}

func doRub(nums []int, index int, resultCache map[int]int) int {
	result, ok := resultCache[index]
	if ok {
		return result
	}
	if index >= len(nums) {
		return 0
	}
	if index == len(nums)-1 {
		resultCache[index] = nums[index]
		return nums[index]
	}

	if index == len(nums) - 2 {
		if nums[index] > nums[index+1] {
			resultCache[index] = nums[index]
			return nums[index]
		}
		resultCache[index] = nums[index+1]
		return nums[index+1]
	}

	firstTrack := nums[index] + doRub(nums, index+2, resultCache)
	secondTrack := doRub(nums, index+1, resultCache)

	if firstTrack > secondTrack {
		resultCache[index] = firstTrack
		return firstTrack
	}
	resultCache[index] = secondTrack
	return secondTrack

}
// @lc code=end

