package leetcode_solutions_golang

//https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	shortest := 2<<31 - 1
	for _, str := range strs {
		if len(str) < shortest {
			shortest = len(str)
		}
	}
	result := ""
	for i := 0; i < shortest; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[j-1][i] {
				return result
			}
		}
		result += string(strs[0][i])
	}
	return result
}
