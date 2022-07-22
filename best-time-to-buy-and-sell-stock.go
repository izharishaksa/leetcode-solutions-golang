package leetcode_solutions_golang

//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	sellPrice := prices[len(prices)-1]
	max := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > sellPrice {
			sellPrice = prices[i]
		}
		if sellPrice-prices[i] > max {
			max = sellPrice - prices[i]
		}
	}
	return max
}
