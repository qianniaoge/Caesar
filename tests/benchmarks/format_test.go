package benchmarks

import (
	"strconv"
	"testing"
)

func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
