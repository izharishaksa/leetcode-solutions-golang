//https://leetcode.com/problems/n-th-tribonacci-number/

package leetcode_solutions_golang

func tribonacci(n int) int {
	dp := make([]int, 38)
	dp[1], dp[2] = 1, 1
	for i := 3; i < 38; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
