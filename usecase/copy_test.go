package usecase

import "testing"

func BenchmarkCopyByCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopyByCopy(xi)
	}
}

func BenchmarkCopyByAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopyByAppend(xi)
	}
}

func BenchmarkCopyByAppendReSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopyByAppendReSlice(xi)
	}
}
