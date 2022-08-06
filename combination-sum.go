package leetcode_solutions_golang

import "container/list"

//https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	type State struct {
		result []int
		total  int
		pos    int
	}

	q := list.New()
	q.PushBack(State{make([]int, 0), 0, 0})

	combinations := make([][]int, 0)
	for q.Len() > 0 {
		el := q.Front()
		curState := el.Value.(State)
		q.Remove(el)
		if curState.total == target {
			combinations = append(combinations, curState.result)
			continue
		}
		for i := curState.pos; i < len(candidates); i++ {
			if curState.total+candidates[i] <= target {
				resultCopy := make([]int, 0)
				for _, j := range curState.result {
					resultCopy = append(resultCopy, j)
				}
				resultCopy = append(resultCopy, candidates[i])
				q.PushBack(State{resultCopy, curState.total + candidates[i], i})
			}
		}
	}

	return combinations
}
