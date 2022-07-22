package leetcode_solutions_golang

//https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/
type ArrayReader interface {
	get(index int) int
}

func search3(reader ArrayReader, target int) int {
	left := 0
	right := 9999
	pivot := (left + right) / 2
	maxVal := 10001
	for right >= left {
		cur := reader.get(pivot)
		if cur == target {
			return pivot
		}
		if cur > maxVal {
			right = pivot - 1
		} else {
			if target > cur {
				right += pivot
				left = pivot + 1
			} else {
				right = pivot - 1
			}
		}
		pivot = (left + right) / 2
	}
	return -1
}
