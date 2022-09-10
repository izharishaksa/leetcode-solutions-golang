package leetcode_solutions_golang

import "testing"

func Test_numberOfWeakCharacters(t *testing.T) {
	type args struct {
		properties [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				properties: [][]int{
					{5, 5},
					{6, 3},
					{3, 6},
				},
			},
			want: 0,
		},
		{
			name: "test case 2",
			args: args{
				properties: [][]int{
					{2, 2},
					{3, 3},
				},
			},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{
				properties: [][]int{
					{1, 5},
					{10, 4},
					{4, 3},
					{4, 2},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeakCharacters(tt.args.properties); got != tt.want {
				t.Errorf("numberOfWeakCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
