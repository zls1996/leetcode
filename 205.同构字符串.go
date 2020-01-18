/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	if (len(s) != len(t)) {
		return false
	}

	indexMap := make(map[byte] byte)

	for index := 0; index < len(s); index++ {
		sVal := s[index]
		indexTValue := t[index]
		tVal, ok := indexMap[sVal]
		if ok {
			if tVal != indexTValue {
				return false
			}
		} else {
			indexMap[sVal] = indexTValue
		}
	}

	resultMap := make(map[byte] byte) 
	for k,v := range indexMap {
		_, ok := resultMap[v]
		if ok {
			return false
		} else {
			resultMap[v] = k
		}
	}

	return true
}

// @lc code=end

