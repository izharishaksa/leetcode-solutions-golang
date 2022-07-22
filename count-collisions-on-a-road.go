package leetcode_solutions_golang

//https://leetcode.com/problems/count-collisions-on-a-road/
func countCollisions(directions string) int {
	arr := []rune(directions)
	total := 0
	size := len(arr)
	for i := 0; i < size; i++ {
		if i < len(arr)-1 && arr[i] == 'R' && arr[i+1] == 'L' {
			arr[i] = 'S'
			arr[i+1] = 'S'
			total += 2
		}
		if i < len(arr)-1 && arr[i] == 'R' && arr[i+1] == 'S' {
			arr[i] = 'S'
			total++
		}
		if i > 0 && arr[i] == 'L' && arr[i-1] == 'S' {
			arr[i] = 'S'
			total++
		}
	}
	for i := 0; i < size; i++ {
		if i < len(arr)-1 && arr[i] == 'R' && arr[i+1] == 'L' {
			arr[i] = 'S'
			arr[i+1] = 'S'
			total += 2
		}
		if i < len(arr)-1 && arr[i] == 'R' && arr[i+1] == 'S' {
			arr[i] = 'S'
			total++
		}
		if i > 0 && arr[i] == 'L' && arr[i-1] == 'S' {
			arr[i] = 'S'
			total++
		}
	}
	for i := size - 1; i >= 0; i-- {
		if i < len(arr)-1 && arr[i] == 'R' && arr[i+1] == 'L' {
			arr[i] = 'S'
			arr[i+1] = 'S'
			total += 2
		}
		if i < len(arr)-1 && arr[i] == 'R' && arr[i+1] == 'S' {
			arr[i] = 'S'
			total++
		}
		if i > 0 && arr[i] == 'L' && arr[i-1] == 'S' {
			arr[i] = 'S'
			total++
		}
	}
	return total
}
