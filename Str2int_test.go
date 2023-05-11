package util4s

import "testing"

func TestStr2Uint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			"test",
			args{
				"1",
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2Uint(tt.args.s); got != tt.want {
				t.Errorf("Str2Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}
