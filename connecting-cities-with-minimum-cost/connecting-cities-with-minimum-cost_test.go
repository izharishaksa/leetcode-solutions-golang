package connecting_cities_with_minimum_cost

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				n: 3,
				connections: [][]int{
					{1, 2, 5},
					{1, 3, 6},
					{2, 3, 1},
				},
			},
			want: 6,
		},
		{
			name: "test case 2",
			args: args{
				n: 4,
				connections: [][]int{
					{1, 2, 3},
					{3, 4, 4},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
