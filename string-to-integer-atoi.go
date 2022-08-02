package leetcode_solutions_golang

import (
	"math"
	"strings"
)

//https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	isNegative := false
	index := -1
	isStarted := false
	if !(s[0] == '-' || (s[0] >= '0' && s[0] <= '9') || s[0] == '+' || s[0] == ' ') {
		isStarted = true
	}
	for i := 0; i < len(s) && !isStarted; i++ {
		if s[i] == '-' {
			isStarted = true
			isNegative = true
			index = i + 1
		}
		if s[i] == '+' {
			isStarted = true
			index = i + 1
		}
		if s[i] >= '0' && s[i] <= '9' {
			index = i
			isStarted = true
		}
	}
	if index < 0 {
		return 0
	}
	result := int64(0)
	for i := index; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int64(s[i]-'0')
			if !isNegative && result > math.MaxInt32 {
				result = math.MaxInt32
			}
			if isNegative && result > math.MaxInt32+1 {
				result = math.MaxInt32 + 1
			}
		} else {
			break
		}
	}
	if isNegative {
		result = -result
	}
	return int(result)
}
