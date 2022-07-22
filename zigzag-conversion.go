package leetcode_solutions_golang

import (
	"sort"
	"strings"
)

type Item struct {
	row   int
	col   int
	value string
}

//https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	curRow := 0
	curCol := 0
	temp := make([]Item, 0)
	isDown := true
	for i := 0; i < len(s); i++ {
		temp = append(temp, Item{
			row:   curRow,
			col:   curCol,
			value: string(s[i]),
		})
		if isDown {
			curRow++
		} else {
			curRow--
			curCol++
		}
		if curRow == numRows-1 || curRow == 0 {
			isDown = !isDown
		}
	}

	sort.Slice(temp, func(i, j int) bool {
		if temp[i].row == temp[j].row {
			if temp[i].col < temp[j].col {
				return true
			}
			return false
		}
		if temp[i].row < temp[j].row {
			return true
		}
		return false
	})

	var res strings.Builder
	for i := 0; i < len(temp); i++ {
		res.WriteString(temp[i].value)
	}

	return res.String()
}
