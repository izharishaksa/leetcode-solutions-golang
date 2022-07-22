package leetcode_solutions_golang

//https://leetcode.com/problems/minimum-bit-flips-to-convert-number/
func minBitFlips(start int, goal int) int {
	if start == goal {
		return 0
	}
	if start > goal {
		start, goal = goal, start
	}
	var count int
	for start != goal {
		if start&1 == goal&1 {
			start >>= 1
			goal >>= 1
		} else {
			start >>= 1
			goal >>= 1
			count++
		}
	}
	return count
}
