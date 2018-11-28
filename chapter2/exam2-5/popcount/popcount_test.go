package popcount_test

import (
	"goBible/chapter2/exam2-5/popcount"
	"testing"
)

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByClearing(0x1234567890ABCDEF)
	}
}
