//https://leetcode.com/problems/numbers-with-same-consecutive-differences/

package leetcode_solutions_golang

import "math"

func numsSameConsecDiff(n int, k int) []int {
	result := make([]int, 0)
	buildNumber(n, k, 0, 0, &result)
	return result
}

func buildNumber(len, diff, curLen int, curNumber int64, result *[]int) {
	if curLen == len {
		*result = append(*result, int(curNumber))
		return
	}
	for i := 0; i < 10; i++ {
		if curLen == 0 && i == 0 {
			continue
		}
		if curLen > 0 && diff != int(math.Abs(float64(i)-float64(curNumber%10))) {
			continue
		}
		buildNumber(len, diff, curLen+1, curNumber*10+int64(i), result)
	}
}
