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

	// TODO 1: comprobar que len(got) es 2

	if len(got) == 2 {

	}

	// TODO 2: comprobar que got[0].ID es "a"

	if got[0].ID == "a" {
		
	}

	// TODO 3: comprobar que got[1].ID es "c"

		if got[1].ID == "c" {
		
	}
}