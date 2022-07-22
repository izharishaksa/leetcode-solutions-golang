package leetcode_solutions_golang

import (
	"container/list"
	"fmt"
)

type State struct {
	n          int
	data       [][]bool
	curRow     int
	col        []bool
	isComplete bool
}

func NewState(n int) *State {
	instance := &State{
		n:          n,
		data:       make([][]bool, n),
		curRow:     0,
		col:        make([]bool, n),
		isComplete: false,
	}
	for i := 0; i < n; i++ {
		instance.data[i] = make([]bool, n)
	}
	return instance
}

func (s *State) Clone() *State {
	instance := &State{
		n:          s.n,
		data:       make([][]bool, s.n),
		curRow:     s.curRow,
		col:        make([]bool, s.n),
		isComplete: s.isComplete,
	}
	for i := 0; i < s.n; i++ {
		instance.data[i] = make([]bool, s.n)
		copy(instance.data[i], s.data[i])
		instance.col[i] = s.col[i]
	}
	return instance
}

func (s *State) GetKey() string {
	return fmt.Sprintf("%v", s.data)
}

func (s *State) String() []string {
	result := make([]string, s.n)
	for i := 0; i < s.n; i++ {
		result[i] = ""
		for j := 0; j < s.n; j++ {
			if s.data[i][j] {
				result[i] += "Q"
			} else {
				result[i] += "."
			}
		}
	}
	return result
}

func (s *State) IncreaseRow() {
	s.curRow++
	if s.curRow == s.n {
		s.isComplete = true
	}
}

func solveNQueens(n int) [][]string {
	q := list.New()
	q.PushBack(NewState(n))
	result := make([][]string, 0)
	generatedKey := make(map[string]*State)
	generatedKey[q.Front().Value.(*State).GetKey()] = q.Front().Value.(*State)

	for q.Len() > 0 {
		el := q.Front()
		q.Remove(el)
		curState := el.Value.(*State)
		if curState.isComplete {
			result = append(result, curState.String())
			continue
		}
		for colIdx := 0; colIdx < n; colIdx++ {
			if curState.col[colIdx] {
				continue
			}
			isOk := true
			for i := 0; i < n; i++ {
				if curState.curRow-i >= 0 && colIdx-i >= 0 && curState.data[curState.curRow-i][colIdx-i] {
					isOk = false
				}
				if curState.curRow-i >= 0 && colIdx+i < n && curState.data[curState.curRow-i][colIdx+i] {
					isOk = false
				}
			}
			if !isOk {
				continue
			}
			newState := curState.Clone()
			newState.data[newState.curRow][colIdx] = true
			newState.IncreaseRow()
			newState.col[colIdx] = true
			key := newState.GetKey()
			if _, ok := generatedKey[key]; !ok {
				generatedKey[key] = newState
				q.PushBack(newState)
			}
		}
	}

	return result
}
