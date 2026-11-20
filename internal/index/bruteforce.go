package index

import (
	"sort"

	"github.com/kev289/minivec/internal/vector"
)

type entry struct {
	id  string
	vec []float32
}
type BruteForce struct {
	dim     int
	entries []entry
	pos     map[string]int
}

var _ Index = (*BruteForce)(nil)

func NewBruteForce(dim int) *BruteForce {
	return &BruteForce{
		dim: dim,
		pos: make(map[string]int),
	}
}

func (b *BruteForce) Add(id string, vec []float32) error {
	if len(vec) != b.dim {
		return vector.ErrDimMismatch
	}

	stored := make([]float32, len(vec))
	copy(stored, vec)

	if i, ok := b.pos[id]; ok {
		b.entries[i].vec = stored
		return nil
	}

	b.pos[id] = len(b.entries)
	b.entries = append(b.entries, entry{id: id, vec: stored})
	return nil
}

func (b *BruteForce) Len() int {
	return len(b.entries)
}

func (b *BruteForce) Search(query []float32, k int) ([]Result, error) {
	if len(query) != b.dim {
		return nil, vector.ErrDimMismatch
	}

	results := make([]Result, 0, len(b.entries))
	for _, e := range b.entries {
		score, err := vector.Cosine(query, e.vec)
		if err != nil {
			return nil, err
		}
		results = append(results, Result{ID: e.id, Score: score})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if k < len(results) {
		results = results[:k]
	}

	return results, nil
}
