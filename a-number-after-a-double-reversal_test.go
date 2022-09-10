package leetcode_solutions_golang

import "testing"

func Test_isSameAfterReversals(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				num: 123,
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				num: 1230,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameAfterReversals(tt.args.num); got != tt.want {
				t.Errorf("isSameAfterReversals() = %v, want %v", got, tt.want)
			}
		})
	}
}
