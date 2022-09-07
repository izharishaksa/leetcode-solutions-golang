//https://leetcode.com/problems/making-a-large-island/

package making_a_large_island

func largestIsland(grid [][]int) int {
	n := len(grid)
	ds := NewDisjointSet(n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i > 0 && grid[i-1][j] == 1 {
					ds.Union(i*n+j, (i-1)*n+j)
				}
				if j > 0 && grid[i][j-1] == 1 {
					ds.Union(i*n+j, i*n+j-1)
				}
			}
		}
	}
	groupSize := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				groupSize[ds.Find(i*n+j)]++
			}
		}
	}
	maxSize := 0
	for _, v := range groupSize {
		maxSize = max(maxSize, v)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				seen := make(map[int]bool)
				size := 1
				if i > 0 && grid[i-1][j] == 1 {
					seen[ds.Find((i-1)*n+j)] = true
				}
				if j > 0 && grid[i][j-1] == 1 {
					seen[ds.Find(i*n+j-1)] = true
				}
				if i < n-1 && grid[i+1][j] == 1 {
					seen[ds.Find((i+1)*n+j)] = true
				}
				if j < n-1 && grid[i][j+1] == 1 {
					seen[ds.Find(i*n+j+1)] = true
				}
				for k := range seen {
					size += groupSize[k]
				}
				maxSize = max(maxSize, size)
			}
		}
	}
	return maxSize
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type DisjointSet struct {
	parent    []int
	groupSize []int
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	groupSize := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		groupSize[i] = 1
	}
	return &DisjointSet{parent, groupSize}
}

func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds *DisjointSet) Union(x, y int) {
	xRoot := ds.Find(x)
	yRoot := ds.Find(y)
	if xRoot == yRoot {
		return
	}
	if ds.groupSize[xRoot] < ds.groupSize[yRoot] {
		ds.parent[xRoot] = yRoot
		ds.groupSize[yRoot] += ds.groupSize[xRoot]
	} else {
		ds.parent[yRoot] = xRoot
		ds.groupSize[xRoot] += ds.groupSize[yRoot]
	}
}
