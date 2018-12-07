package isprime_test

import (
	"testing"

	"goBible/chapter3/isprime"
)

func BenchmarkIsPrimeSimle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isprime.IsPrimeSimple(400000)
	}
}

func BenchmarkIsPrimeImprove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isprime.IsPrimeImprove(400000)
	}
}

func BenchmarkIsPrimeBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isprime.IsPrimeBest(400000)
	}
}
