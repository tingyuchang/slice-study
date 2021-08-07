package usecase

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "Basic Reverse",
			args: args{
				[]interface{}{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
			},
			want: []interface{}{
				9, 8, 7, 6, 5, 4, 3, 2, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse2(t *testing.T) {
	type args struct {
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "Switch Reverse",
			args: args{
				[]interface{}{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
			},
			want: []interface{}{
				9, 8, 7, 6, 5, 4, 3, 2, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse2(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffling(t *testing.T) {
	type args struct {
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			"Shuffle Test",
			args{
				[]interface{}{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
			},
			[]interface{}{
				1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shuffling(tt.args.x); reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffling() = %v, dont want %v", got, tt.want)
			}
		})
	}
}

func TestBatch(t *testing.T) {
	type args struct {
		x         []interface{}
		batchSize int
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "Test for Batch slice",
			args: args{
				x: []interface{}{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				},
				batchSize: 3,
			},
			want: [][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Batch(tt.args.x, tt.args.batchSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Batch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatch2(t *testing.T) {
	type args struct {
		x         []interface{}
		batchSize int
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "Test for Batch slice",
			args: args{
				x: []interface{}{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				},
				batchSize: 3,
			},
			want: [][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Batch2(tt.args.x, tt.args.batchSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Batch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func BenchmarkReverse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse2([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func BenchmarkBatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Batch([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	}
}

func BenchmarkBatch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Batch2([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	}
}

func TestDedupalicate(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test for de-duplicate",
			args: args{
				[]int{3, 2, 3, 2, 4, 2, 3, 1, 2, 3, 4, 4, 2, 5, 6, 7, 2, 6, 4},
			},
			want: []int{
				1, 2, 3, 4, 5, 6, 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dedupalicate(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dedupalicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDedupalicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dedupalicate([]int{3, 2, 3, 2, 4, 2, 3, 1, 2, 3, 4, 4, 2, 5, 6, 7, 2, 6, 4})
	}
}

func TestMoveToFront(t *testing.T) {
	type args struct {
		needle   string
		haystack []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test move to front",
			args: args{
				needle:   "c",
				haystack: []string{"a", "b", "c", "d", "e"},
			},
			want: []string{"c", "a", "b", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveToFront(tt.args.needle, tt.args.haystack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveToFront() = %v, want %v", got, tt.want)
			}
		})
	}
}
