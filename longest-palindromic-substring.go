package leetcode_solutions_golang

//https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	maxResult := ""
	for i := 0; i < len(s); i++ {
		result := string(s[i])
		left := i
		right := i
		for {
			if left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
				result = string(s[left-1]) + result + string(s[right+1])
			} else {
				break
			}
			left--
			right++
		}
		if len(result) > len(maxResult) {
			maxResult = result
		}
		if i+1 < len(s) && s[i] == s[i+1] {
			result = string(s[i]) + string(s[i+1])
			left = i
			right = i + 1
			for {
				if left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
					result = string(s[left-1]) + result + string(s[right+1])
				} else {
					break
				}
				left--
				right++
			}
			if len(result) > len(maxResult) {
				maxResult = result
			}
		}
	}
	return maxResult
}
