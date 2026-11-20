package index

import (
	"sort"
	"runtime"
	"sync"

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

	n := len(b.entries)
	results := make([]Result, n)

	workers := runtime.NumCPU()
	chunk := (n + workers - 1) / workers
	errs := make([]error, workers)

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		start := w * chunk
		end := min(start+chunk, n)
		if start >= end {
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := start; i < end;  i ++ {
				score, err := vector.Cosine(query, b.entries[i].vec)
			if err != nil {
				errs[w] = err
				return 
			}
			results[i] = Result{ID: b.entries[i].id, Score: score }
			}
		}()
	}
	wg.Wait()

	for _, err :=  range errs {
		if err != nil{
			return nil, err
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if k < len(results) {
		results = results[:k]
	}
	return results, nil	
}
