//https://leetcode.com/problems/connecting-cities-with-minimum-cost/

package connecting_cities_with_minimum_cost

import "sort"

func minimumCost(n int, connections [][]int) int {
	if n == 1 {
		return 0
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	ds := NewDisjointSet(n)
	total := 0
	for _, c := range connections {
		if ds.Find(c[0]) != ds.Find(c[1]) {
			ds.Union(c[0], c[1])
			total += c[2]
			n--
		}
	}
	if n > 1 {
		return -1
	}
	return total
}

type DisjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &DisjointSet{parent, rank}
}

func (d *DisjointSet) Find(x int) int {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
}

func (d *DisjointSet) Union(x, y int) {
	xRoot := d.Find(x)
	yRoot := d.Find(y)
	if xRoot == yRoot {
		return
	}
	if d.rank[xRoot] < d.rank[yRoot] {
		d.parent[xRoot] = yRoot
	} else if d.rank[xRoot] > d.rank[yRoot] {
		d.parent[yRoot] = xRoot
	} else {
		d.parent[yRoot] = xRoot
		d.rank[xRoot]++
	}
}
