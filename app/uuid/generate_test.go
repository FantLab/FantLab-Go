package uuid

import (
	"testing"
)

func Benchmark_Generate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateNow()
	}
}

func Benchmark_enerateParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GenerateNow()
		}
	})
}
