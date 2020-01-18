/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxProfit := 0
	buyPrice := prices[0]
	sellPrice := buyPrice

	for index := 1; index < len(prices); index++ {
		price := prices[index]
		if price > sellPrice {
			sellPrice = price
			if sellPrice-buyPrice > maxProfit {
				maxProfit = sellPrice-buyPrice
			}
		}  else if price < buyPrice {
			sellPrice = price
			buyPrice = price
		}
	}

	if sellPrice-buyPrice > maxProfit {
		maxProfit = sellPrice-buyPrice
	}

	return maxProfit
	
}
// @lc code=end

