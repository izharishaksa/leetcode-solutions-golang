package combination_sum_ii

import (
	"fmt"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test case 1",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !compareSlice(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}

		})
	}
}

func compareSlice(got [][]int, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}
	for _, w := range want {
		if !contains(got, w) {
			return false
		}
	}
	return true
}

func contains(got [][]int, w []int) bool {
	for _, g := range got {
		if fmt.Sprintf("%v", g) == fmt.Sprintf("%v", w) {
			return true
		}
	}
	return false
}
