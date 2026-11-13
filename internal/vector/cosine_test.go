package vector

import (
	"errors"
	"math"
	"testing"
)

func TestCosine(t *testing.T) {
	tests := []struct {
		name    string
		a, b    []float32
		want    float32
		wantErr error
	}{
		{"identicos", []float32{1, 2, 3}, []float32{1, 2, 3}, 1, nil},
		{"ortogonales", []float32{1, 0}, []float32{0, 1}, 0, nil},
		{"opuestos", []float32{1, 2}, []float32{-1, -2}, -1, nil},
		{"caso real", []float32{1, 2, 3}, []float32{4, 5, 6}, 0.97463185, nil},
		{"longitudes distintas", []float32{1, 2}, []float32{1, 2, 3}, 0, ErrDimMismatch},
		{"vector cero", []float32{0, 0}, []float32{1, 2}, 0, ErrZeroVector},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cosine(tt.a, tt.b)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if math.Abs(float64(got-tt.want)) > 1e-6 {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}