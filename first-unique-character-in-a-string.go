//https://leetcode.com/problems/first-unique-character-in-a-string/

package leetcode_solutions_golang

func firstUniqChar(s string) int {
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if count[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
