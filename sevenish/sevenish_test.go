package sevenish

import "testing"

func TestFindSevenishNumber(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "2",
			args: args{n: 2},
			want: 7,
		},
		{
			name: "3",
			args: args{n: 3},
			want: 8,
		},
		{
			name: "4",
			args: args{n: 4},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSevenishNumber(tt.args.n); got != tt.want {
				t.Errorf("FindSevenishNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
