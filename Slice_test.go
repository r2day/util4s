package util4s

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicateStr(t *testing.T) {
	type args struct {
		strSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test for dup",
			args{
				[]string{
					"a", "b", "b",
				},
			},
			[]string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicateStr(tt.args.strSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicateStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
