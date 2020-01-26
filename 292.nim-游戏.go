/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start
func canWinNim(n int) bool {
	return canWin(n, true)
}

func canWin(n int, myRound bool) bool {
	if n >= 1 && n <= 3 {
		return myRound
	}

	if n % 4 == 0 {
		return !myRound
	}

	return myRound
	
}
// @lc code=end

