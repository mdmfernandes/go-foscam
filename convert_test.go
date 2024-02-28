package foscam

import "testing"

func Test_b2u(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "true",
			args: args{true},
			want: 1,
		},
		{
			name: "false",
			args: args{false},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b2u(tt.args.b); got != tt.want {
				t.Errorf("b2u() = %v, want %v", got, tt.want)
			}
		})
	}
}
