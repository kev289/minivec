# minivec

A small vector database written in Go. I built it to learn Go's concurrency
model and to practice designing a system in layers. It is a learning project,
not production software.

## What it does

`minivec` stores float vectors and finds the ones most similar to a query
vector using cosine similarity.

## Status

Milestone 1 is done: cosine similarity, the `Index` interface and a
brute-force index.

## Roadmap

- [x] **M1**: cosine similarity, `Index` interface, brute-force index
- [ ] **M2**: concurrent search (goroutines, `context`) with benchmarks
- [ ] **M3**: safe concurrent writes (`sync.RWMutex`, race detector)
- [ ] **M4**: ingestion pipeline (worker pool, rate limiter, `Embedder` interface)
- [ ] **M5**: approximate index (simplified HNSW) compared against brute force
- [ ] **M6**: persistence (write-ahead log and snapshots)
- [ ] **M7**: HTTP API and a small RAG chatbot as a demo client

## Project layout

```
cmd/minivec/        entry point
internal/vector/    vector math (cosine similarity)
internal/index/     Index interface and implementations
docs/               design notes
```

## Design notes

- `Index` is a small interface (`Add`, `Search`, `Len`). Implementations are
  interchangeable, so a brute-force index and an approximate one can be
  benchmarked against each other.
- The embedding model will sit behind an `Embedder` interface, so the core
  does not depend on any provider.

## Development

```bash
go test ./...
go vet ./...
```