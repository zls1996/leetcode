/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	indexMap := make(map[int] int)

	for index := 0; index < len(nums); index++ {
		val := nums[index]
		_, ok := indexMap[val]
		if ok {
			return true
		} else {
			indexMap[val] = index
		}
	}

	return false
}
// @lc code=end

