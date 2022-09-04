package leetcode_solutions_golang

import "testing"

func Test_findSubarrays(t *testing.T) {
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
				nums: []int{4, 2, 4},
			},
			want: true,
		},
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("findSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
