package word_break_ii

import (
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test case 1",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{"cats and dog", "cat sand dog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !compareSlice(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareSlice(got []string, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for _, w := range want {
		if !contains(got, w) {
			return false
		}
	}
	return true
}

func contains(got []string, w string) bool {
	for _, g := range got {
		if g == w {
			return true
		}
	}
	return false
}
