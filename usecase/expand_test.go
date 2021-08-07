package usecase

import "testing"

func BenchmarkInsertByAppend(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Insert(xi, 5, 99, 98, 97, 96, 95, 94, 93, 92)
	}
}

func BenchmarkInsertByCopy(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Insert2(xi, 5, 99, 98, 97, 96, 95, 94, 93, 92)
	}
}
