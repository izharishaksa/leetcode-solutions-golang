package leetcode_solutions_golang

//https://leetcode.com/problems/valid-palindrome-ii/
func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	counter := 0
	for left <= right {
		if counter > 1 {
			break
		}
		if s[left] != s[right] {
			counter++
			right--
		} else {
			left++
			right--
		}
	}
	if counter < 2 {
		return true
	}
	left = 0
	right = len(s) - 1
	counter = 0
	for left <= right {
		if counter > 1 {
			return false
		}
		if s[left] != s[right] {
			counter++
			left++
		} else {
			left++
			right--
		}
	}
	return true
}
