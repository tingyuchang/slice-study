package usecase

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		x    []interface{}
		keep func(interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "filter odd",
			args: args{
				x: []interface{}{
					1,2,3,4,5,6,7,8,9,10,11,12,
				},
				keep: func(i interface{}) bool {
					if i.(int) %2 == 0 {
						return false
					}
					return true
				},
			},
			want: []interface{}{
				1,3,5,7,9,11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.x, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
