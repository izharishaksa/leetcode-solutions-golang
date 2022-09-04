//https://leetcode.com/problems/strictly-palindromic-number/

package strictly_palindromic_number

import "math/big"

func isStrictlyPalindromic(n int) bool {
	x := n - 2
	if x > 62 {
		x = 62
	}
	for i := 2; i <= x; i++ {
		s := big.NewInt(int64(n)).Text(i)
		if !isPalindromic(s) {
			return false
		}
	}
	return true
}

func isPalindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
