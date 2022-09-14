package leetcode_solutions_golang

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				s: "MCDXLIV",
			},
			want: 1444,
		},
		{
			name: "test case 2",
			args: args{
				s: "CMXCIX",
			},
			want: 999,
		},
		{
			name: "test case 3",
			args: args{
				s: "CXI",
			},
			want: 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
