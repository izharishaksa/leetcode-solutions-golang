package leetcode_solutions_golang

//https://leetcode.com/problems/find-first-palindromic-string-in-the-array/
func isPalindrome2(word string) bool {
	left := 0
	right := len(word) - 1
	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome2(word) {
			return word
		}
	}
	return ""
}
