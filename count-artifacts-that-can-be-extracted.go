package leetcode_solutions_golang

//https://leetcode.com/problems/count-artifacts-that-can-be-extracted/
func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	nOccupied := make(map[int]int, 0)
	counter := 1
	for _, artifact := range artifacts {
		nOccupied[counter] = 0
		for i := artifact[0]; i <= artifact[2]; i++ {
			for j := artifact[1]; j <= artifact[3]; j++ {
				arr[i][j] = counter
				nOccupied[counter]++
			}
		}
		counter++
	}
	for _, d := range dig {
		artifactId := arr[d[0]][d[1]]
		nOccupied[artifactId]--
	}
	result := 0
	for _, val := range nOccupied {
		if val == 0 {
			result++
		}
	}
	return result
}
