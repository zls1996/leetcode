/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	Sums []int
	Nums []int
}


func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray {
			Sums: [] int {},
			Nums: nums,
		}
	}
	sums := make([] int, len(nums))
	sums[0] = nums[0]
	for index := 1; index < len(nums); index++ {
		sums[index] = sums[index-1] + nums[index]
	}
	return NumArray {
		Sums: sums,
		Nums: nums,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.Sums[j]
	} else {
		return this.Sums[j] - this.Sums[i-1]
	}
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

