//https://leetcode.com/problems/remove-duplicate-letters/

package leetcode_solutions_golang

func removeDuplicateLetters(s string) string {
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	isAdded := make([]bool, 26)
	result := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		for len(result) > 0 && result[len(result)-1] > s[i] && count[result[len(result)-1]-'a'] > 1 && !isAdded[s[i]-'a'] {
			count[result[len(result)-1]-'a']--
			isAdded[result[len(result)-1]-'a'] = false
			result = result[:len(result)-1]
		}
		if !isAdded[s[i]-'a'] {
			result = append(result, s[i])
			isAdded[s[i]-'a'] = true
		} else {
			count[s[i]-'a']--
		}
	}
	return string(result)
}
