/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
    if len(s) <= 0 {
		return true
	}

	begin := 0
	end := len(s) - 1

	for end >= begin {
		first := s[begin]
		last := s[end]
		if !isLetterOrNumber(first) {
			begin++
			continue
		}

		if !isLetterOrNumber(last) {
			end--
			continue
		}

		if !sameIgnoreCase(first , last) {
			return false
		}
		begin++
		end--
	}

	return true
}

func isLetterOrNumber(by byte) bool {
	if by >= '0' && by <= '9' {
		return true
	} 
	if by >= 'A' && by <= 'Z' {
		return true
	}

	if by >= 'a' && by <= 'z' {
		return true
	}

	return false
}

func sameIgnoreCase(first byte, second byte) bool {
	if first == second {
		return true
	}
	if first < 'A' || second < 'A' {
		return false
	}
	var val int
    if first > second {
        val = int(first-second)
    } else {
        val = int(second-first)
    }
	if val == 32 {
		return true
	}

	return false
	
}
// @lc code=end

