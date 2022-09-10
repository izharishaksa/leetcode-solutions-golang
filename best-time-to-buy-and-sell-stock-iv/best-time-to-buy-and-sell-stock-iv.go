//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/

package best_time_to_buy_and_sell_stock_iv

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) <= 1 {
		return 0
	}
	dp := make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][k][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
