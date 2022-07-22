package leetcode_solutions_golang

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	table := make([]int, 10001)
	for i := 0; i < len(coins); i++ {
		if coins[i] < 10001 {
			table[coins[i]] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for _, j := range coins {
			if i > j && table[i-j] != 0 && table[i-j]+table[j] < min {
				min = table[i-j] + table[j]
			}
		}
		if min != math.MaxInt32 && table[i] == 0 {
			table[i] = min
		}
	}

	if int(table[amount]) == 0 {
		return -1
	}
	return table[amount]
}
