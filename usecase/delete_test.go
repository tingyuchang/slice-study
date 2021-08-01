package usecase

import "testing"

func BenchmarkDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Delete(5, xi)
	}
}

func BenchmarkDeleteByCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DeleteByCopy(5, xi)
	}
}
