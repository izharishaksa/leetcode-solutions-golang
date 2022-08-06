//https://leetcode.com/problems/counting-bits/
package leetcode_solutions_golang

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}

//11=1011, 1011 >> 1 = 101 = result[5] = 2
//11=1011, 1011 & 0001 = 1
//result[11] = result[5] + (11 & 1) = 2 + 1 = 3
