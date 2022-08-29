package leetcode_solutions_golang

import (
	"fmt"
)

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	orderByRow, err := buildTopologicalSort(k, rowConditions)
	if err != nil {
		return [][]int{}
	}
	orderByCol, err := buildTopologicalSort(k, colConditions)
	if err != nil {
		return [][]int{}
	}

	matrix := make([][]int, k)
	for i := 0; i < k; i++ {
		matrix[i] = make([]int, k)
		for j := 0; j < k; j++ {
			if orderByRow[i] == orderByCol[j] {
				matrix[i][j] = orderByRow[i]
			}
		}
	}

	return matrix
}

func buildTopologicalSort(nNode int, edge [][]int) ([]int, error) {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for i := 0; i < len(edge); i++ {
		graph[edge[i][0]] = append(graph[edge[i][0]], edge[i][1])
		inDegree[edge[i][1]]++
	}

	order := make([]int, 0)
	for {
		next := -1
		isCompleted := true
		for i := 1; i <= nNode; i++ {
			if inDegree[i] >= 0 {
				isCompleted = false
			}
			if inDegree[i] == 0 {
				next = i
				break
			}
		}
		if isCompleted {
			return order, nil
		}
		if next == -1 {
			return nil, fmt.Errorf("cycle detected")
		}

		order = append(order, next)
		inDegree[next] = -1
		for _, v := range graph[next] {
			inDegree[v]--
		}
	}
}
