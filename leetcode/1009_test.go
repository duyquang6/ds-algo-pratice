package leetcode

import "testing"

func Test_bitwiseComplement(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				N: 5,
			},
			want: 2,
		},
		{
			args: args{
				N: 7,
			},
			want: 0,
		},
		{
			args: args{
				N: 10,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplement(tt.args.N); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
