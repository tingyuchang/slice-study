package main

import "testing"

type Object struct {
	Value int
}

func BenchmarkWithPointer(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		xo := make([]*Object, 100)
		for i := 0; i < 100; i++ {
			xo = append(xo, &Object{i})
		}
	}
}

func BenchmarkWithoutPointer(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		xo := make([]Object, 100)
		for i := 0; i < 100; i++ {
			xo = append(xo, Object{i})
		}
	}
}
