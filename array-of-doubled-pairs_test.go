package leetcode_solutions_golang

import "testing"

func Test_canReorderDoubled(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				arr: []int{3, 1, 3, 6},
			},
			want: false,
		},
		{
			name: "test case 2",
			args: args{
				arr: []int{2, 1, 2, 6},
			},
			want: false,
		},
		{
			name: "test case 3",
			args: args{
				arr: []int{4, -2, 2, -4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReorderDoubled(tt.args.arr); got != tt.want {
				t.Errorf("canReorderDoubled() = %v, want %v", got, tt.want)
			}
		})
	}
}
