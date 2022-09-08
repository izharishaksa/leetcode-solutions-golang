package number_of_islands_ii

import (
	"reflect"
	"testing"
)

func Test_numIslands2(t *testing.T) {
	type args struct {
		height    int
		width     int
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				height: 3,
				width:  3,
				positions: [][]int{
					{0, 0},
					{1, 0},
					{1, 2},
					{1, 1},
				},
			},
			want: []int{1, 1, 2, 1},
		},
		{
			name: "test case 2",
			args: args{
				height: 2,
				width:  2,
				positions: [][]int{
					{0, 0},
					{1, 1},
					{0, 1},
				},
			},
			want: []int{1, 2, 1},
		},
		{
			name: "test case 3",
			args: args{
				height: 3,
				width:  3,
				positions: [][]int{
					{0, 0},
					{0, 1},
					{1, 2},
					{1, 2},
				},
			},
			want: []int{1, 1, 2, 2},
		},
		{
			name: "test case 4",
			args: args{
				height: 3,
				width:  3,
				positions: [][]int{
					{0, 1},
					{0, 0},
					{1, 2},
					{1, 2},
				},
			},
			want: []int{1, 1, 2, 2},
		},
		{
			name: "test case 5",
			args: args{
				height: 3,
				width:  3,
				positions: [][]int{
					{0, 0},
					{1, 0},
					{0, 1},
					{1, 1},
				},
			},
			want: []int{1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands2(tt.args.height, tt.args.width, tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numIslands2() = %v, want %v", got, tt.want)
			}
		})
	}
}
