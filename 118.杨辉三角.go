/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
    if numRows <= 0 {
		return make([][] int, 0)
	}
	result := make([][] int ,0, numRows)

	for index := 0; index < numRows; index++ {
		columnLength := index + 1
		array := make([] int, columnLength)
		for i := 0; i < columnLength; i++ {
			if i == 0 || i == columnLength-1 {
				array[i] = 1
				continue
			}
			if index != 0 {
				lastArray := result[index-1]
				array[i] = lastArray[i-1] + lastArray[i]
			}
		}
		result = append(result, array)
	}

	return result
}
// @lc code=end

