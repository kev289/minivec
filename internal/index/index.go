package index

type Result struct {
	ID    string
	Score float32
}

type Index interface {
	Add(id string, vec []float32) error
	Search(query []float32, k int) ([]Result, error)
	Len() int
}
