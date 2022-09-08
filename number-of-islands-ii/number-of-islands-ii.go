//https://leetcode.com/problems/number-of-islands-ii/

package number_of_islands_ii

func numIslands2(height int, width int, positions [][]int) []int {
	ds := NewDisjointSet(height * width)
	seen := make(map[int]bool)
	result := make([]int, len(positions))
	totalGroup := 0
	for index, pos := range positions {
		row := pos[0]
		col := pos[1]
		curId := row*width + col
		if seen[curId] {
			result[index] = totalGroup
			continue
		}
		seen[curId] = true
		neighborGroups := make(map[int]bool, 0)
		if row > 0 && seen[(row-1)*width+col] {
			neighborGroups[ds.Find((row-1)*width+col)] = true
		}
		if col > 0 && seen[row*width+col-1] {
			neighborGroups[ds.Find(row*width+col-1)] = true
		}
		if row < height-1 && seen[(row+1)*width+col] {
			neighborGroups[ds.Find((row+1)*width+col)] = true
		}
		if col < width-1 && seen[row*width+col+1] {
			neighborGroups[ds.Find(row*width+col+1)] = true
		}
		curGroup := ds.Find(curId)

		if len(neighborGroups) == 0 {
			totalGroup++
		} else {
			for group := range neighborGroups {
				ds.Union(curGroup, group)
			}
			mergedGroups := make(map[int]bool, 0)
			for group := range neighborGroups {
				mergedGroups[ds.Find(group)] = true
			}
			totalGroup -= len(neighborGroups) - len(mergedGroups)
		}
		result[index] = totalGroup
	}
	return result
}

type DisjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSet{parent, rank}
}

func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds *DisjointSet) Union(x int, y int) {
	xRoot := ds.Find(x)
	yRoot := ds.Find(y)
	if xRoot == yRoot {
		return
	}
	if ds.rank[xRoot] < ds.rank[yRoot] {
		ds.parent[xRoot] = yRoot
	} else if ds.rank[xRoot] > ds.rank[yRoot] {
		ds.parent[yRoot] = xRoot
	} else {
		ds.parent[yRoot] = xRoot
		ds.rank[xRoot]++
	}
}
