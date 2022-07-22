package leetcode_solutions_golang

//https://leetcode.com/problems/largest-3-same-digit-number-in-string/
func largestGoodInteger(num string) string {
	max := ' '
	result := ""
	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] && num[i] > byte(max) {
			max = rune(num[i])
			result = string(num[i]) + string(num[i+1]) + string(num[i+2])
		}
	}
	return result
}
