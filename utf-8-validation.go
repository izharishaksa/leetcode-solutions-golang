//https://leetcode.com/problems/utf-8-validation/

package leetcode_solutions_golang

func validUtf8(data []int) bool {
	count := 0
	for _, v := range data {
		if count == 0 {
			if v>>5 == 0b110 {
				count = 1
			} else if v>>4 == 0b1110 {
				count = 2
			} else if v>>3 == 0b11110 {
				count = 3
			} else if v>>7 != 0 {
				return false
			}
		} else {
			if v>>6 != 0b10 {
				return false
			}
			count--
		}
	}
	return count == 0
}
