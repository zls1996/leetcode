/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}

	indexMap := make(map[int] int)

	for index := 0; index < len(nums); index++ {
		val := nums[index]
		lastIndex, ok := indexMap[val]
		if ok {
			if index - lastIndex <= k {
				return true
			}
		}
		indexMap[val] = index
	}
	return false
}
// @lc code=end

