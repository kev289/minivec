package index

import "github.com/kev289/minivec/internal/vector"

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
