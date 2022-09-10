//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
