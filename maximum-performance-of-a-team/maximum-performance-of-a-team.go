//https://leetcode.com/problems/maximum-performance-of-a-team/

package maximum_performance_of_a_team

import (
	"container/heap"
	"sort"
)

type Engineer struct {
	Speed      int
	Efficiency int
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	engs := make([]Engineer, n)
	for i := 0; i < n; i++ {
		engs[i] = Engineer{speed[i], efficiency[i]}
	}
	sort.Slice(engs, func(i, j int) bool {
		return engs[i].Efficiency > engs[j].Efficiency
	})
	speedHeap := &IntHeap{}
	heap.Init(speedHeap)
	speedSum := 0
	ret := 0
	for _, eng := range engs {
		heap.Push(speedHeap, eng.Speed)
		speedSum += eng.Speed
		if speedHeap.Len() > k {
			speedSum -= heap.Pop(speedHeap).(int)
		}
		ret = max(ret, speedSum*eng.Efficiency)
	}
	return ret % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
