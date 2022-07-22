package leetcode_solutions_golang

//https://leetcode.com/problems/reverse-string/
func revString(s []byte, left, right int) {
	if left >= right {
		return
	}
	s[left] ^= s[right]
	s[right] ^= s[left]
	s[left] ^= s[right]
	left++
	right--
	revString(s, left, right)
}

func reverseString(s []byte) {
	revString(s, 0, len(s)-1)
}
