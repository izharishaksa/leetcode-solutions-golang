package leetcode_solutions_golang

//https://leetcode.com/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(-n) == n
}
