package min_cost_to_connect_all_points

import "testing"

func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				points: [][]int{
					{0, 0},
					{2, 2},
					{3, 10},
					{5, 2},
					{7, 0},
				},
			},
			want: 20,
		},
		{
			name: "test case 2",
			args: args{
				points: [][]int{
					{0, 0},
				},
			},
			want: 0,
		},
		{
			name: "test case 3",
			args: args{
				points: [][]int{
					{2, -3}, {-17, -8}, {13, 8}, {-17, -15},
				},
			},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
