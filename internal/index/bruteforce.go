package index

import (
	"sort"

	"github.com/kev289/minivec/internal/vector"
)

type BruteForce struct {
	dim  int
	vecs map[string][]float32
}

func NewBruteForce(dim int) *BruteForce {
	return &BruteForce{
		dim:  dim,
		vecs: make(map[string][]float32),
	}
}

func (b *BruteForce) Add(id string, vec []float32) error {
	if len(vec) != b.dim {
		return vector.ErrDimMismatch
	}

	stored := make([]float32, len(vec))
	copy(stored, vec)
	b.vecs[id] = stored
	return nil
}

func (b *BruteForce) Len() int {
	return len(b.vecs)
}

func (b *BruteForce) Search(query []float32, k int) ([]Result, error) {
	if len(query) != b.dim {
		return nil, vector.ErrDimMismatch
	}

	results := make([]Result, 0, len(b.vecs))
	for id, vec := range b.vecs {
		score, err := vector.Cosine(query, vec)
		if err != nil {
			return nil, err
		}
		results = append(results, Result{ID: id, Score: score})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if k < len(results) {
		results = results[:k]
	}

	return results, nil
}
