package leetcode_solutions_golang

import "testing"

func Test_kIncreasing(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{12, 6, 12, 6, 14, 2, 13, 17, 3, 8, 11, 7, 4, 11, 18, 8, 8, 3},
				k:   1,
			},
			want: 12,
		},
		{
			name: "test2",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
				k:   1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kIncreasing(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("kIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
