package best_time_to_buy_and_sell_stock_iv

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 0",
			args: args{
				k:      0,
				prices: []int{},
			},
			want: 0,
		},
		{
			name: "test case 1",
			args: args{
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
