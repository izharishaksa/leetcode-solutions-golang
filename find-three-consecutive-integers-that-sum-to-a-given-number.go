//https://leetcode.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/

package leetcode_solutions_golang

func sumOfThree(num int64) []int64 {
	mid := num / 3
	if ((mid - 1) + mid + (mid + 1)) == num {
		return []int64{mid - 1, mid, mid + 1}
	}
	return []int64{}
}
