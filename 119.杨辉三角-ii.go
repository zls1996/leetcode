/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
    if rowIndex <= 0 {
		return []int{1}
	}

	result := make([] int, rowIndex+1)
	begin := 0
	end := len(result) -1 

	for end >= begin {
		if begin == 0 {
			result[begin] = 1
		} else {
			queryMap := make(map[QueryKey]int)
			result[begin] = getLast(rowIndex-1, begin - 1, queryMap) + getLast(rowIndex-1, begin, queryMap)
		}
		result[end] = result[begin]
		end--
		begin++
	}

	return result
}

func getLast(rowIndex int, columnIndex int, queryMap map[QueryKey]int) int {
	queryKey := QueryKey{
		RowIndex: rowIndex,
		ColumnIndex: columnIndex,
	}

	result , ok := queryMap[queryKey]
	if ok {
		return result
	}

	if columnIndex == 0 || rowIndex == columnIndex {
		return 1
	}

	val := getLast(rowIndex-1, columnIndex - 1, queryMap) + getLast(rowIndex-1, columnIndex, queryMap)
	queryMap[queryKey] = val
	return val
}

type QueryKey struct {
	RowIndex int
	ColumnIndex int
}
// @lc code=end

