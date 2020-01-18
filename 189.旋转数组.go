/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
    doRotate(nums, 0, k)
}

func doRotate(nums []int, begin int , end int) {
	if (begin == end) {
		return
	}
	final := nums[len(nums)-1]
	for index := len(nums)-2; index >= 0 ; index-- {
		val := nums[index]
		nums[index+1] = val
	}
	nums[0] = final

	doRotate(nums, begin+1, end)
}
// @lc code=end

