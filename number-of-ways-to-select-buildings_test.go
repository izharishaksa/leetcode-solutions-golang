package leetcode_solutions_golang

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				s: "001101",
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				s: "101010",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.s); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
