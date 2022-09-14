package leetcode_solutions_golang

import (
	"reflect"
	"testing"
)

func Test_sumOfThree(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test case 1",
			args: args{
				num: 6,
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfThree(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
