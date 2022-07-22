package leetcode_solutions_golang

//https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	var x = [100]int{0, 1, 2}
	for i := 3; i <= n; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	return x[n]
}
