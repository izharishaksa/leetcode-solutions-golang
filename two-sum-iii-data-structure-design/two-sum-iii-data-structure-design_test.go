package two_sum_iii_data_structure_design

import "testing"

func TestTwoSum(t1 *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				number: 4,
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				number: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t *testing.T) {
			twoSum := Constructor()
			twoSum.Add(1)
			twoSum.Add(1)
			twoSum.Add(3)
			if got := twoSum.Find(tt.args.number); got != tt.want {
				t.Errorf("TwoSum.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
