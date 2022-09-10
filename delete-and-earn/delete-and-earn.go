//https://leetcode.com/problems/delete-and-earn/

package delete_and_earn

import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	keys := make([]int, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	dp := make([]int, len(keys))
	dp[0] = keys[0] * count[keys[0]]
	for i := 1; i < len(keys); i++ {
		if keys[i]-keys[i-1] == 1 {
			if i > 1 {
				dp[i] = max(dp[i-1], dp[i-2]+keys[i]*count[keys[i]])
			} else {
				dp[i] = max(dp[i-1], keys[i]*count[keys[i]])
			}
		} else {
			dp[i] = dp[i-1] + keys[i]*count[keys[i]]
		}
	}

	return dp[len(dp)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
