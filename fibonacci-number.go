package leetcode_solutions_golang

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	res := make([]int, n+1)
	res[1] = 1
	res[2] = 1
	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n]
}
