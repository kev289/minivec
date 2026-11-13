package vector

import "errors"

var ErrDimMismatch = errors.New("vector: dimension mismatch")

func Cosine(a, b []float32) (float32, error) {
	if len(a) != len(b) {
		return 0, ErrDimMismatch
	}
	return 0, nil
}