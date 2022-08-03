package leetcode_solutions_golang

import "strings"

//https://leetcode.com/problems/text-justification/
func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	index := 0
	curLine := make([]string, 0)
	curLen := 0
	for index < len(words) {
		requiredLen := len(words[index])
		if len(curLine) > 0 {
			requiredLen++
		}
		if curLen+requiredLen <= maxWidth {
			curLine = append(curLine, words[index])
			curLen += requiredLen
		} else {
			space := maxWidth - curLen + (len(curLine) - 1)
			line := curLine[0]
			if len(curLine) == 1 {
				for j := 0; j < space; j++ {
					line += " "
				}
			} else {
				eachSpace := space / (len(curLine) - 1)
				resSpace := space % (len(curLine) - 1)

				for i := 1; i < len(curLine); i++ {
					for j := 0; j < eachSpace; j++ {
						line += " "
					}
					if resSpace > 0 {
						line += " "
						resSpace--
					}
					line += curLine[i]
				}
			}
			result = append(result, line)

			curLine = make([]string, 0)
			curLine = append(curLine, words[index])
			curLen = len(words[index])
		}
		index++
	}
	if len(curLine) > 0 {
		line := strings.Join(curLine, " ")
		for i := len(line); i < maxWidth; i++ {
			line += " "
		}
		result = append(result, line)
	}
	return result
}
