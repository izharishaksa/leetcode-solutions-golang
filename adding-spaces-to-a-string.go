package leetcode_solutions_golang

import "bytes"

//https://leetcode.com/problems/adding-spaces-to-a-string/
func addSpaces(s string, spaces []int) string {
	isSpaced := make(map[int]bool)
	for _, index := range spaces {
		isSpaced[index] = true
	}
	buffer := new(bytes.Buffer)
	buffer.WriteString("")

	for i := 0; i < len(s); i++ {
		_, isRequired := isSpaced[i]
		if isRequired {
			buffer.WriteString(" ")
		}
		buffer.WriteString(string(s[i]))
	}

	return buffer.String()
}
