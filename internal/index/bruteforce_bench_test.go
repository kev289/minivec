package index

import (
	"fmt"
	"math/rand"
	"testing"
)

func randomVector(r *rand.Rand, dim int) []float32 {
	v := make([]float32, dim)
	for i := range v {
		v[i] = r.Float32()*2 - 1
	}
	return v
}

func BenchmarkBruteForceSearch(b *testing.B) {
	const (
		n   = 10_000
		dim = 128
	)
	r := rand.New(rand.NewSource(42))

	idx := NewBruteForce(dim)
	for i := 0; i < n; i++ {
		id := fmt.Sprintf("v%d", i)
		if err := idx.Add(id, randomVector(r, dim)); err != nil {
			b.Fatal(err)
		}
	}
	query := randomVector(r, dim)

	b.ResetTimer()
	for b.Loop() {
		if _, err := idx.Search(query, 10); err != nil {
			b.Fatal(err)
		}
	}
}
