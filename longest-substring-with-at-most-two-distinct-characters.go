package leetcode_solutions_golang

import (
	"math"
)

//https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
func lengthOfLongestSubstringTwoDistinct(s string) int {
	max := 0
	cur := 0
	last := 0
	appeared := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		_, ok := appeared[s[i]]
		if !ok && len(appeared) > 1 {
			prev := s[i-1]
			for k, _ := range appeared {
				if k != prev {
					delete(appeared, k)
				}
			}
			cur -= last
		}
		appeared[s[i]]++
		if i > 0 && s[i] != s[i-1] {
			last = i
		}
		cur++
		max = int(math.Max(float64(max), float64(cur)))
	}
	return max
}
