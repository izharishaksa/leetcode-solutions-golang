package leetcode_solutions_golang

import "testing"

func Test_bestHand(t *testing.T) {
	type args struct {
		ranks []int
		suits []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				ranks: []int{1, 1, 1, 2, 2},
				suits: []byte{'S', 'H', 'C', 'D', 'C'},
			},
			want: "Three of a Kind",
		},
		{
			name: "test case 2",
			args: args{
				ranks: []int{1, 1, 2, 3},
				suits: []byte{'S', 'C', 'D', 'C'},
			},
			want: "Pair",
		},
		{
			name: "test case 3",
			args: args{
				ranks: []int{1, 2, 3, 4, 5},
				suits: []byte{'S', 'H', 'C', 'D', 'C'},
			},
			want: "High Card",
		},
		{
			name: "test case 4",
			args: args{
				ranks: []int{1, 2, 3, 4, 5},
				suits: []byte{'S', 'S', 'S', 'S', 'S'},
			},
			want: "Flush",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestHand(tt.args.ranks, tt.args.suits); got != tt.want {
				t.Errorf("bestHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
