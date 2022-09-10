package leetcode_solutions_golang

import "testing"

func Test_garbageCollection(t *testing.T) {
	type args struct {
		garbage []string
		travel  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				garbage: []string{"MP", "PGG", "GM"},
				travel:  []int{1, 2, 3},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := garbageCollection(tt.args.garbage, tt.args.travel); got != tt.want {
				t.Errorf("garbageCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
