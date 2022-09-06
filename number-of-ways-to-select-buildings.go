//https://leetcode.com/problems/number-of-ways-to-select-buildings/
package leetcode_solutions_golang

func numberOfWays(s string) int64 {
	size := len(s)
	total := int64(0)
	byZero := make([]int, size)
	sumZero := make([]int, size)
	curZero := 0

	byOne := make([]int, size)
	sumOne := make([]int, size)
	curOne := 0

	for i := range s {
		if s[i] == '1' {
			curOne++
			byZero[i] = curZero
		} else {
			curZero++
			byOne[i] = curOne
		}
		sumZero[i] = curZero
		sumOne[i] = curOne
	}

	for i := 0; i < size; i++ {
		total += int64(byZero[i] * (sumZero[size-1] - sumZero[i]))
		total += int64(byOne[i] * (sumOne[size-1] - sumOne[i]))
	}

	return total
}
