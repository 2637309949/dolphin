package reflect

import (
	"reflect"
	"testing"
)

func TestOrOne(t *testing.T) {
	type args struct {
		arr []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "TestOrOne",
			args: struct {
				arr []interface{}
			}{
				arr: []interface{}{0, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrOne(tt.args.arr...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
