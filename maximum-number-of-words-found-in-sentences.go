package leetcode_solutions_golang

import "strings"

//https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
func mostWordsFound(sentences []string) int {
	max := 0
	for _, str := range sentences {
		cur := strings.Split(str, " ")
		if len(cur) > max {
			max = len(cur)
		}
	}
	return max
}
