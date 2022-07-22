package leetcode_solutions_golang

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x *= -1
	}
	res := float64(0)
	for x > 0 {
		res *= 10
		res += float64(x % 10)
		x /= 10
	}
	if isNegative {
		res *= -1
	}
	maxInt := 2147483647
	minInt := -2147483648
	if res > float64(maxInt) || res < float64(minInt) {
		return 0
	}
	return int(res)
}
