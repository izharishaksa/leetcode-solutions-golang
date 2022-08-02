package leetcode_solutions_golang

//https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'C':
			if i < len(s)-1 && s[i+1] == 'D' {
				result += 400
				i++
			} else if i < len(s)-1 && s[i+1] == 'M' {
				result += 900
				i++
			} else {
				result += 100
			}
		case 'X':
			if i < len(s)-1 && s[i+1] == 'L' {
				result += 40
				i++
			} else if i < len(s)-1 && s[i+1] == 'C' {
				result += 90
				i++
			} else {
				result += 10
			}
		case 'I':
			if i < len(s)-1 && s[i+1] == 'V' {
				result += 4
				i++
			} else if i < len(s)-1 && s[i+1] == 'X' {
				result += 9
				i++
			} else {
				result += 1
			}
		default:
			result += roman[string(s[i])]
		}
	}
	return result
}
