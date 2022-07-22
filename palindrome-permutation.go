package leetcode_solutions_golang

//https://leetcode.com/problems/palindrome-permutation/
func canPermutePalindrome(s string) bool {
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	totalNegative := 0
	for _, v := range count {
		if v%2 == 1 {
			totalNegative++
		}
	}
	if totalNegative < 2 {
		return true
	}
	return false
}
