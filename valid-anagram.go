package leetcode_solutions_golang

func isAnagram(s string, t string) bool {
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		count[t[i]-'a']--
	}
	for _, i := range count {
		if i != 0 {
			return false
		}
	}
	return true
}
