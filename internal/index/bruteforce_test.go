package index

import (
	"errors"
	"testing"

	"github.com/kev289/minivec/internal/vector"
)

func TestBruteForceAdd(t *testing.T) {
	idx := NewBruteForce(3)

	err := idx.Add("a", []float32{1, 2, 3})
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	if idx.Len() != 1 {
		t.Errorf("Len = %d, want 1", idx.Len())
	}

	err = idx.Add("b", []float32{1, 2})
	if !errors.Is(err, vector.ErrDimMismatch) {
		t.Fatalf("Add returned unexpected error: %v", err)
	}
}