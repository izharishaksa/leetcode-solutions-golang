//https://leetcode.com/problems/construct-smallest-number-from-di-string/

package construct_smallest_number_from_di_string

import (
	"container/list"
	"fmt"
)

type State struct {
	pos    int
	num    int
	result string
	count  []int
}

func (s *State) IsValid() bool {
	for i := 0; i < len(s.count); i++ {
		if s.count[i] > 1 {
			return false
		}
	}
	return true
}

func smallestNumber(pattern string) string {
	q := list.New()
	for i := 1; i < 10; i++ {
		s := State{
			pos:    0,
			num:    i,
			result: fmt.Sprintf("%d", i),
			count:  make([]int, 10),
		}
		s.count[i] = 1
		q.PushBack(s)
	}

	for q.Len() > 0 {
		cur := q.Front().Value.(State)
		q.Remove(q.Front())
		if cur.pos == len(pattern) {
			return cur.result
		}
		if pattern[cur.pos] == 'I' {
			for i := cur.num + 1; i < 10; i++ {
				s := State{
					pos:    cur.pos + 1,
					num:    i,
					result: cur.result + fmt.Sprintf("%d", i),
					count:  make([]int, 10),
				}
				copy(s.count, cur.count)
				s.count[i]++
				if s.IsValid() {
					q.PushBack(s)
				}
			}
		} else {
			for i := 1; i < cur.num; i++ {
				s := State{
					pos:    cur.pos + 1,
					num:    i,
					result: cur.result + fmt.Sprintf("%d", i),
					count:  make([]int, 10),
				}
				copy(s.count, cur.count)
				s.count[i]++
				if s.IsValid() {
					q.PushBack(s)
				}
			}
		}
	}

	return ""
}
