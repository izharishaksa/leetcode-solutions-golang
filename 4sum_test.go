package leetcode_solutions_golang

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test case 1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{
				{-1, 0, 0, 1},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); compareSlices(got, tt.want) == false {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareSlices(got [][]int, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := range got {
		if contains(want, got[i]) == false {
			return false
		}
	}

	return true
}

func contains(want [][]int, ints []int) bool {
	for i := range want {
		if reflect.DeepEqual(want[i], ints) {
			return true
		}
	}

	return false
}
