package leetcode_solutions_golang

import (
	"reflect"
	"testing"
)

func Test_findOriginalArray(t *testing.T) {
	type args struct {
		changed []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				changed: []int{1, 3, 4, 2, 6, 8},
			},
			want: []int{4, 3, 1},
		},
		{
			name: "test case 2",
			args: args{
				changed: []int{6, 3, 0, 1},
			},
			want: []int{},
		},
		{
			name: "test case 3",
			args: args{
				changed: []int{1},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOriginalArray(tt.args.changed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOriginalArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
