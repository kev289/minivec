package vector

import (
	"errors"
	"math"
)

var ErrDimMismatch = errors.New("vector: dimension mismatch")

// Cosine devuelve la similitud coseno entre a y b.
func Cosine(a, b []float32) (float32, error) {
	if len(a) != len(b) {
		return 0, ErrDimMismatch
	}

	var dot, normA, normB float64
	for i := range a {
		// TODO: acumular dot, normA y normB
		// pista: convierte con float64(a[i])
	}

	// TODO: decidir qué pasa si normA o normB es 0
	// TODO: devolver dot / (math.Sqrt(normA) * math.Sqrt(normB)) convertido a float32
}