package leetcode_solutions_golang

//https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/
func getDescentPeriods(prices []int) int64 {
	dp := [100001]int64{}
	for index, _ := range prices {
		dp[index] = 1
	}
	total := dp[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1]-1 {
			dp[i] += dp[i-1]
		}
		total += dp[i]
	}
	return total
}
