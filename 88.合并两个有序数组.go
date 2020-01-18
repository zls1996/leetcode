/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	if len(nums1) == 0 {
		return
	}

	if len(nums2) == 0 {
		return
	}
	lastIndex := m+n-1
	mIndex := m-1
	nIndex := n-1

	for mIndex >= 0 && nIndex >= 0 {
		if nums1[mIndex] > nums2[nIndex] {
			nums1[lastIndex] = nums1[mIndex]
			mIndex--
		} else {
			nums1[lastIndex] = nums2[nIndex]
			nIndex--
		}
		lastIndex--
	}

	if mIndex < 0 {
		for nIndex >= 0 {
			nums1[lastIndex] = nums2[nIndex]
			lastIndex--
			nIndex--
		}	
		return
	}

	if nIndex < 0 {
		for mIndex >= 0 {
			nums1[lastIndex] = nums1[mIndex]
			lastIndex--
			mIndex--
		}	
		return
	}
	
}
// @lc code=end

