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

func TestBruteForceSearch(t *testing.T) {
	idx := NewBruteForce(2)

	if err := idx.Add("a", []float32{1, 0}); err != nil {
		t.Fatalf("Add a: %v", err)
	}
	if err := idx.Add("b", []float32{0, 1}); err != nil {
		t.Fatalf("Add b: %v", err)
	}
	if err := idx.Add("c", []float32{1, 1}); err != nil {
		t.Fatalf("Add c: %v", err)
	}

	got, err := idx.Search([]float32{1, 0}, 2)
	if err != nil {
		t.Fatalf("Search returned error: %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("len(got) = %d, want 2", len(got))
	}

	if got[0].ID != "a" {
		t.Errorf("got[0].ID = %q, want a", got[0].ID)
	}

	if got[1].ID != "c" {
		t.Errorf("got[1].ID = %q, want c", got[1].ID)
	}
}
