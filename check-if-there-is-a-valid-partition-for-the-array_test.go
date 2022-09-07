package leetcode_solutions_golang

import "testing"

func Test_validPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{1, 1, 1, 2},
			},
			want: false,
		},
		{
			name: "test case 2",
			args: args{
				nums: []int{4, 4, 4, 5, 6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPartition(tt.args.nums); got != tt.want {
				t.Errorf("validPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
