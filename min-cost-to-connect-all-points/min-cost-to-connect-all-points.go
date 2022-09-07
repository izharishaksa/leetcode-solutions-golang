//https://leetcode.com/problems/min-cost-to-connect-all-points/

package min_cost_to_connect_all_points

import "container/heap"

type Edge struct {
	From, To, Distance int
}

type Edges []Edge

func (e Edges) Len() int {
	return len(e)
}

func (e Edges) Less(i, j int) bool {
	return e[i].Distance < e[j].Distance
}

func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *Edges) Push(x interface{}) {
	*e = append(*e, x.(Edge))
}

func (e *Edges) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func minCostConnectPoints(points [][]int) int {
	if len(points) == 1 {
		return 0
	}
	pq := &Edges{}
	heap.Init(pq)
	added := make([]bool, len(points))
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			heap.Push(pq, Edge{
				From:     i,
				To:       j,
				Distance: abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1]),
			})
		}
	}
	first := heap.Pop(pq).(Edge)
	added[first.From] = true
	added[first.To] = true
	cost := first.Distance
	isProgressing := true
	for isProgressing {
		temp := make([]Edge, 0)
		isProgressing = false
		for pq.Len() > 0 {
			edge := heap.Pop(pq).(Edge)
			if added[edge.From] && added[edge.To] {
				continue
			}
			isProgressing = true
			if added[edge.From] || added[edge.To] {
				added[edge.From] = true
				added[edge.To] = true
				cost += edge.Distance
				break
			} else {
				temp = append(temp, edge)
			}
		}
		for e := 0; e < len(temp); e++ {
			heap.Push(pq, temp[e])
		}
	}

	return cost
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
