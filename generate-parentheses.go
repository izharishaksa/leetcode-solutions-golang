package leetcode_solutions_golang

import "container/list"

//https://leetcode.com/problems/generate-parentheses/
type ParenthesesState struct {
	len    int
	open   int
	close  int
	result string
}

func (s ParenthesesState) IsComplete() bool {
	if s.open == s.len && s.close == s.len {
		return true
	}
	return false
}

func generateParenthesis(n int) []string {
	q := list.New()
	q.PushBack(ParenthesesState{n, 0, 0, ""})
	result := make([]string, 0)
	for q.Len() > 0 {
		el := q.Front()
		curState := el.Value.(ParenthesesState)
		q.Remove(el)
		if curState.IsComplete() {
			result = append(result, curState.result)
		} else {
			if curState.open >= curState.close && curState.open < curState.len {
				q.PushBack(ParenthesesState{curState.len, curState.open + 1, curState.close, curState.result + "("})
			}
			if curState.close < curState.open && curState.close < curState.len {
				q.PushBack(ParenthesesState{curState.len, curState.open, curState.close + 1, curState.result + ")"})
			}
		}
	}
	return result
}
