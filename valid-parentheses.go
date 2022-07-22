package leetcode_solutions_golang

import "container/list"

//https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	stack := list.New()
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack.PushFront(c)
		case ')', ']', '}':
			if stack.Len() == 0 {
				return false
			}
			el := stack.Front()
			stack.Remove(el)
			cur := el.Value.(uint8)
			if (c == ')' && cur != '(') || (c == ']' && cur != '[') || (c == '}' && cur != '{') {
				return false
			}
		}
	}
	if stack.Len() > 0 {
		return false
	}
	return true
}
