package strictly_palindromic_number

import "testing"

func Test_isStrictlyPalindromic(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				n: 9,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStrictlyPalindromic(tt.args.n); got != tt.want {
				t.Errorf("isStrictlyPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
