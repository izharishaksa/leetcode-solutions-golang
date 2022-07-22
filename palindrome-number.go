package leetcode_solutions_golang

//https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	res := 0
	for x > 0 {
		res *= 10
		res += x % 10
		x /= 10
	}
	if res == original {
		return true
	}
	return false
}
