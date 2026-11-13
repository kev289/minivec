package vector

import "testing"

func TestDotTemporal(t *testing.T) {
	got, _ := Cosine([]float32{1, 2, 3}, []float32{4, 5, 6})
	t.Log(got)
}