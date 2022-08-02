package leetcode_solutions_golang

//https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	for i := 0; i < len(roman); i++ {
		for num >= values[i] {
			result += roman[i]
			num -= values[i]
		}
	}
	return result
}
