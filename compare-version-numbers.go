package leetcode_solutions_golang

import (
	"strconv"
	"strings"
)

//https://leetcode.com/problems/compare-version-numbers/
func compareVersion(version1 string, version2 string) int {
	v1Slice := strings.Split(version1, ".")
	v2Slice := strings.Split(version2, ".")

	if len(v1Slice) < len(v2Slice) {
		diff := len(v2Slice) - len(v1Slice)
		for i := 0; i < diff; i++ {
			v1Slice = append(v1Slice, "0")
		}
	} else if len(v1Slice) > len(v2Slice) {
		diff := len(v1Slice) - len(v2Slice)
		for i := 0; i < diff; i++ {
			v2Slice = append(v2Slice, "0")
		}
	}

	size := len(v1Slice)
	for i := 0; i < size; i++ {
		a, _ := strconv.Atoi(v1Slice[i])
		b, _ := strconv.Atoi(v2Slice[i])
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
	}

	return 0
}
