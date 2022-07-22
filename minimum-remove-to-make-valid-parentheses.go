package leetcode_solutions_golang

import "container/list"

//https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
func minRemoveToMakeValid(s string) string {
	skipped := make(map[int]bool)
	stack := list.New()
	stackPos := list.New()
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.PushFront(s[i])
			stackPos.PushFront(i)
		} else if s[i] == ')' {
			last := stack.Front()
			lastPost := stackPos.Front()
			if last != nil && last.Value.(byte) == '(' {
				stack.Remove(last)
				stackPos.Remove(lastPost)
			} else {
				skipped[i] = true
			}
		}
	}
	for e := stackPos.Front(); e != nil; e = e.Next() {
		skipped[e.Value.(int)] = true
	}
	for i := 0; i < len(s); i++ {
		if !skipped[i] {
			result += string(s[i])
		}
	}

	return result
}
