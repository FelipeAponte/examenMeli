package mutant

import (
	"testing"
)

// test differents DNA threads
func TestIsMutant(t *testing.T) {
	for _, test := range dnaThreads {
		result := IsMutant(test.thread)
		if result != test.expected {
			t.Errorf("got %v, wanted %v", result, test.expected)
		}
	}
}

// completely crosses a 120x120 matrix that has no coincidence
func BenchmarkDna120x120(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsMutant(DnaThread120x120)
	}
}
