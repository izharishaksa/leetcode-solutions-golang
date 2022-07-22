package leetcode_solutions_golang

//https://leetcode.com/problems/maximum-split-of-positive-even-integers/
func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum == 2 || finalSum == 4 {
		return []int64{finalSum}
	}
	if finalSum%2 != 0 {
		return []int64{}
	}
	var result []int64
	for i := 2; ; i += 2 {
		finalSum -= int64(i)
		next := int64((i + 2) * 2)
		result = append(result, int64(i))
		if finalSum-next <= 0 {
			result = append(result, finalSum)
			break
		}
	}
	return result
}
