//https://leetcode.com/problems/count-special-integers/

package leetcode_solutions_golang

import "fmt"

func count(number, cur, len, size int, digits []int, result *int) {
	if len < size {
		temp := cur * 10
		for i := 0; i < 10; i++ {
			if digits[i] == 0 {
				newNumber := temp + i
				if newNumber > 0 && newNumber <= number {
					*result++
					digits[i] = 1
					count(number, temp+i, len+1, size, digits, result)
					digits[i] = 0
				}
			}
		}
	}
}

func countSpecialNumbers(n int) int {
	s := fmt.Sprintf("%d", n)
	size := len(s)

	result := 0
	count(n, 0, 0, size, make([]int, 10), &result)

	return result
}
