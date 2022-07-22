package leetcode_solutions_golang

import "strings"

//https://leetcode.com/problems/valid-palindrome/
func isPalindrome3(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		if !((s[left] >= 'a' && s[left] <= 'z') || (s[left] >= '0' && s[left] <= '9')) {
			left++
			continue
		}
		if !((s[right] >= 'a' && s[right] <= 'z') || (s[right] >= '0' && s[right] <= '9')) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
