package vector

import (
	"errors"
	"math"
)

var (
	ErrDimMismatch = errors.New("vector: dimension mismatch")
	ErrZeroVector  = errors.New("vector: zero vector")
)

func Cosine(a, b []float32) (float32, error) {
	if len(a) != len(b) {
		return 0, ErrDimMismatch
	}

	var dot, normA, normB float64
	for i := range a {
		x := float64(a[i])
		y := float64(b[i])
		dot += x * y
		normA += x * x
		normB += y * y
	}

	if normA == 0 || normB == 0 {
		return 0, ErrZeroVector
	}

	result := dot / (math.Sqrt(normA) * math.Sqrt(normB))
	return float32(result), nil
}