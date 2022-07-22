package leetcode_solutions_golang

import "math"

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	cur := 0
	max := 0
	appeared := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		index, ok := appeared[s[i]]
		if !ok {
			cur++
			appeared[s[i]] = i
		} else {
			for k, v := range appeared {
				if v <= index {
					cur--
					delete(appeared, k)
				}
			}
			appeared[s[i]] = i
			cur++
		}
		max = int(math.Max(float64(max), float64(cur)))
	}

	return max
}
