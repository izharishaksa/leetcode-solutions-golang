//https://leetcode.com/problems/removing-stars-from-a-string/

package leetcode_solutions_golang

func removeStars(s string) string {
	res := ""
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			count++
		} else if count > 0 {
			count--
		} else {
			res = string([]byte{s[i]}) + res
		}
	}
	return res
}
