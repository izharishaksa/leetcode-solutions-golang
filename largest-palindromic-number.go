//https://leetcode.com/problems/largest-palindromic-number/

package leetcode_solutions_golang

import (
	"fmt"
	"strings"
)

func largestPalindromic(num string) string {
	count := make([]int, 10)
	for i := 0; i < len(num); i++ {
		count[num[i]-'0']++
	}
	mid := ""
	for i := 9; i >= 0; i-- {
		if count[i]%2 == 1 {
			mid = fmt.Sprintf("%d", i)
			break
		}
	}
	var left, right strings.Builder
	isOk := false
	for i := 9; i >= 0; i-- {
		for j := 0; j < count[i]/2; j++ {
			if i > 0 {
				isOk = true
			}
			if i == 0 && !isOk {
				continue
			}
			left.WriteString(fmt.Sprintf("%d", i))
			right.WriteString(fmt.Sprintf("%d", i))
		}
	}
	if left.Len() == 0 && mid == "" && right.Len() == 0 {
		return "0"
	}
	left.WriteString(mid)
	temp := right.String()
	for i := len(temp) - 1; i >= 0; i-- {
		left.WriteString(string(temp[i]))
	}
	return left.String()
}
