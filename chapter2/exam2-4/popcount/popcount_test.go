package popcount_test

import (
	"goBible/chapter2/exam2-4/popcount"
	"testing"
)

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByShifting(0x1234567890ABCDEF)
	}
}
