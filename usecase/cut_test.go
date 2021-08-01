package usecase

import "testing"

func BenchmarkCut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cut(2,8, xi)
	}
}

func BenchmarkCutByCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CutByCopy(2, 8, xi)
	}
}