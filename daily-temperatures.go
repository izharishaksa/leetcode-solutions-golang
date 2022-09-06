//https://leetcode.com/problems/daily-temperatures/

package leetcode_solutions_golang

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	temp := make([]int, 101)
	temp[temperatures[len(temperatures)-1]] = len(temperatures) - 1
	for i := len(temperatures) - 2; i >= 0; i-- {
		for j := temperatures[i] + 1; j < 101; j++ {
			if temp[j] > 0 && (result[i] == 0 || temp[j]-i < result[i]) {
				result[i] = temp[j] - i
			}
		}
		temp[temperatures[i]] = i
	}
	return result
}
