package maximum_rows_covered_by_columns

import "testing"

func Test_maximumRows(t *testing.T) {
	type args struct {
		mat  [][]int
		cols int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				mat: [][]int{
					{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1},
				},
				cols: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRows(tt.args.mat, tt.args.cols); got != tt.want {
				t.Errorf("maximumRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
