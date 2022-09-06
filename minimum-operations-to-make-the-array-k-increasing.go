//https://leetcode.com/problems/minimum-operations-to-make-the-array-k-increasing/

package leetcode_solutions_golang

func kIncreasing(arr []int, k int) int {
	byK := make([][]int, k)
	for i := 0; i < len(arr); i++ {
		byK[i%k] = append(byK[i%k], arr[i])
	}
	total := 0
	for i := 0; i < k; i++ {
		lis := make([]int, 0)
		lis = append(lis, byK[i][0])
		for j := 1; j < len(byK[i]); j++ {
			if byK[i][j] > lis[len(lis)-1] {
				lis = append(lis, byK[i][j])
			} else {
				index := upperBound(lis, byK[i][j])
				if index == len(lis) {
					lis = append(lis, byK[i][j])
				} else {
					lis[index] = byK[i][j]
				}
			}
		}
		total += len(byK[i]) - len(lis)
	}
	return total
}

func upperBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		if target < arr[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
