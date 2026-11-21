notas de rendimiento (bruteforce)

Mi PC: Ryzen 7 3700X, 8 núcleos / 16 hilos.
Benchmark: 10.000 vectores, dim 128, k=10.
Comando: `go test -run '^$' -bench=BruteForceSearch -benchmem -cpu 1,16 ./internal/index`

| qué probé | ms/op | allocs/op |
|---|---|---|
| sort.Slice (reflexión) | ~1,96 | 24 |
| slices.SortFunc (genérico) | ~1,42 | 20 |
| sin sort (con `if false`), 1 hilo | ~0,275 | |
| sin sort (con `if false`), 16 hilos | ~0,303 | |

Con SortFunc da lo mismo con 1 que con 16 hilos.

lo que aprendí:
- el sort se lleva ~1,15 ms de los 1,42 ms, más del 80%. Lo lento era ordenar, no calcular cosenos.
- repartir en 16 goroutines con 10k vectores no gana nada: cada una hace ~625 vectores y lanzarlas cuesta lo mismo que lo que ahorran.
- pprof (go tool pprof) mostraba sort.Slice y reflectlite.Swapper como lo más pesado. Cambiar a slices.SortFunc quitó la reflexión y bajó de 1,96 a 1,42.
- los 4,12 ms de antes los medí en otra máquina (8 hilos) con el código secuencial, así que no se pueden comparar con estos.

pendiente:
- top-k con container/heap en vez de ordenar todo
- benchmark con 1k, 10k y 100k para ver desde cuándo las goroutines sí ganan
