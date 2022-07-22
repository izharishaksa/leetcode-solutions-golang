package leetcode_solutions_golang

import "sort"

//https://leetcode.com/problems/boats-to-save-people/
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	size := len(people) - 1
	total := 0
	for i, j := 0, size; i <= j; j, total = j-1, total+1 {
		if people[j]+people[i] <= limit {
			i++
		}
	}
	return total
}
