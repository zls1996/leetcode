/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	return commonPrefix(strs[0], 1, strs)

}

func commonPrefix(prefix string, index int, strs []string) string {
	target := strs[index]
	var resultStr string
	if len(prefix) > len(target) {
		resultStr = doGetCommonPrefix(target, prefix)
	} else {
		resultStr = doGetCommonPrefix(prefix, target)
	}

	if (index >= len(strs) -1) {
		return resultStr
	}

	index++
	return commonPrefix(resultStr, index, strs)
}

func doGetCommonPrefix(prefix, target string) string {
	if prefix == "" || target == "" {
		return ""
	}

	var prefixBytes []byte = []byte(prefix)
	var targetBytes []byte = []byte(target)

	
	if prefixBytes[0] != targetBytes[0] {
		return ""
	}

	var offset int = 1;

	for i := 1; i < len(prefixBytes); i++ {
		if prefixBytes[i] == targetBytes[i] {
			offset++;
		} else {
			break
		}
	}

	return string(targetBytes[0 : offset])

}
// @lc code=end

