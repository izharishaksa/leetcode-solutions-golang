//https://leetcode.com/problems/a-number-after-a-double-reversal/

package leetcode_solutions_golang

func isSameAfterReversals(num int) bool {
	origin := num
	result := 0
	for num > 0 {
		result *= 10
		result += num % 10
		num /= 10
	}
	result2 := 0
	for result > 0 {
		result2 *= 10
		result2 += result % 10
		result /= 10
	}
	if result2 == origin {
		return true
	}
	return false
}
