//https://leetcode.com/problems/two-sum-iii-data-structure-design/

package two_sum_iii_data_structure_design

type TwoSum struct {
	number map[int]bool
	sum    map[int]bool
}

func Constructor() TwoSum {
	return TwoSum{
		number: make(map[int]bool),
		sum:    make(map[int]bool),
	}
}

func (t *TwoSum) Add(number int) {
	for k, _ := range t.number {
		t.sum[number+k] = true
	}
	if t.number[number] {
		t.sum[number+number] = true
	}
	t.number[number] = true
}

func (t *TwoSum) Find(value int) bool {
	if t.sum[value] {
		return true
	}
	return false
}
