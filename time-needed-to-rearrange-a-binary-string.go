//https://leetcode.com/problems/time-needed-to-rearrange-a-binary-string/

package leetcode_solutions_golang

import "strings"

func secondsToRemoveOccurrences(s string) int {
	total := 0
	for {
		newStr := strings.Replace(s, "01", "10", -1)
		if newStr == s {
			return total
		}
		total++
		s = newStr
	}
}
