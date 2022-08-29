package leetcode_solutions_golang

import (
	"reflect"
	"testing"
)

func Test_buildMatrix(t *testing.T) {
	type args struct {
		k             int
		rowConditions [][]int
		colConditions [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				k:             3,
				rowConditions: [][]int{{1, 2}, {3, 2}},
				colConditions: [][]int{{2, 1}, {3, 2}},
			},
			want: [][]int{{0, 0, 1}, {3, 0, 0}, {0, 2, 0}},
		},
		{
			name: "test2",
			args: args{
				k:             3,
				rowConditions: [][]int{{1, 2}, {2, 3}, {2, 3}, {2, 1}},
				colConditions: [][]int{{2, 1}},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildMatrix(tt.args.k, tt.args.rowConditions, tt.args.colConditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
