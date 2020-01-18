/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxProfit := 0
	buyPrice := prices[0]

	for index := 1; index < len(prices); index++ {
		price := prices[index]
		if price > buyPrice {
			maxProfit += (price-buyPrice)
		} 
		buyPrice = price
	}

	return maxProfit
}
// @lc code=end

