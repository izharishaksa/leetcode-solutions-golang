//https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/

package leetcode_solutions_golang

import (
	"container/list"
	"sort"
)

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]
		}
		return properties[i][0] < properties[j][0]
	})

	stack := list.New()
	for _, i := range properties {
		for stack.Len() > 0 {
			el := stack.Front()
			front := el.Value.([]int)
			if front[0] < i[0] && front[1] < i[1] {
				stack.Remove(el)
			} else {
				break
			}
		}
		stack.PushFront(i)
	}

	return len(properties) - stack.Len()
}
