package usecase

import "testing"

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insert(xi, 5, 99,98,97,96,95,94,93,92)
	}
}

func BenchmarkInsert2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insert2(xi, 5, 99,98,97,96,95,94,93,92)
	}
}