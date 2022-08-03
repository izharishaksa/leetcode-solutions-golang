package leetcode_solutions_golang

//https://leetcode.com/problems/letter-combinations-of-a-phone-number/
var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	generateCombination(0, "", digits, &result)
	return result
}

func generateCombination(index int, cur, digits string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, cur)
		return
	}
	alphabet, _ := phoneMap[string(digits[index])]
	for _, v := range alphabet {
		generateCombination(index+1, cur+string(v), digits, result)
	}
}
