package leetcode_solutions_golang

import (
	"container/list"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := list.New()
	for _, data := range tokens {
		number, err := strconv.Atoi(data)
		if err == nil {
			stack.PushFront(number)
			continue
		}
		elA := stack.Front()
		stack.Remove(elA)
		elB := stack.Front()
		stack.Remove(elB)
		a := elA.Value.(int)
		b := elB.Value.(int)

		result := 0
		switch data {
		case "+":
			result = a + b
			break
		case "-":
			result = b - a
			break
		case "*":
			result = a * b
			break
		case "/":
			result = b / a
			break
		}
		stack.PushFront(result)
	}
	return stack.Front().Value.(int)
}
