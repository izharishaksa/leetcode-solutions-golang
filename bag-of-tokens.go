//https://leetcode.com/problems/bag-of-tokens/

package leetcode_solutions_golang

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	score, maxScore := 0, 0
	left, right := 0, len(tokens)-1
	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			left++
			score++
			if score > maxScore {
				maxScore = score
			}
		} else if score > 0 {
			power += tokens[right]
			right--
			score--
		} else {
			break
		}
	}

	return maxScore
}
