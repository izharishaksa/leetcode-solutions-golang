package leetcode_solutions_golang

//https://leetcode.com/problems/factorial-trailing-zeroes/
func trailingZeroes(n int) int {
	ret := 0
	for n >= 5 {
		n /= 5
		ret += n
	}
	return ret
}
