package leetcode_solutions_golang

import "fmt"

func groupAnagrams(strs []string) [][]string {
	mapped := make(map[string][]string, 0)
	for i := 0; i < len(strs); i++ {
		count := make([]byte, 26)
		for j := 0; j < len(strs[i]); j++ {
			count[strs[i][j]-'a']++
		}
		key := fmt.Sprintf("%+v", count)
		cur, isExist := mapped[key]
		if !isExist {
			cur = make([]string, 0)
		}
		cur = append(cur, strs[i])
		mapped[key] = cur
	}

	result := make([][]string, 0)
	for _, v := range mapped {
		result = append(result, v)
	}

	return result
}
