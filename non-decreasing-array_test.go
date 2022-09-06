package leetcode_solutions_golang

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test0",
			args: args{
				nums: []int{1, 4, 2, 5},
			},
			want: true,
		},
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				nums: []int{3, 4, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
