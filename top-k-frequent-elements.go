package leetcode_solutions_golang

import "container/heap"

type Data struct {
	Val  int
	Freq int
}

type FreqHeap []Data

func (h FreqHeap) Len() int {
	return len(h)
}

func (h FreqHeap) Less(i, j int) bool {
	if h[i].Freq > h[j].Freq {
		return true
	}
	return false
}

func (h FreqHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FreqHeap) Push(x interface{}) {
	*h = append(*h, x.(Data))
}

func (h *FreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, 0)
	for _, i := range nums {
		cur, _ := freq[i]
		freq[i] = cur + 1
	}

	fh := &FreqHeap{}
	heap.Init(fh)
	for k, v := range freq {
		heap.Push(fh, Data{Val: k, Freq: v})
	}

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		data := heap.Pop(fh).(Data)
		result = append(result, data.Val)
	}

	return result
}
