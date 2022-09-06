package leetcode_solutions_golang

import "testing"

func Test_removeDuplicateLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "bcabc",
			},
			want: "abc",
		},
		{
			name: "test2",
			args: args{
				s: "cbacdcbc",
			},
			want: "acdb",
		},
		{
			name: "test3",
			args: args{
				s: "bbcaac",
			},
			want: "bac",
		},
		{
			name: "test4",
			args: args{
				s: "abacb",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
