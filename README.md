## Go Cache Benchmarks

Benchmarks of in-memory cache libraries for Go with expiration/TTL support.

## Run Test

    go get
    go test -bench=. -benchmem

## Results

goos: darwin
goarch: amd64
pkg: github.com/yogesh-desai/go-cache-benchmark
BenchmarkKodingCache/Set-4               1000000              1732 ns/op             376 B/op          2 allocs/op
BenchmarkKodingCache/Get-4               2000000               619 ns/op               7 B/op          0 allocs/op
BenchmarkHashicorpLRU/Set-4              3000000               411 ns/op             102 B/op          4 allocs/op
BenchmarkHashicorpLRU/Get-4             20000000                82.6 ns/op             8 B/op          0 allocs/op
BenchmarkCache2Go/Set-4                  1000000              2561 ns/op             466 B/op         10 allocs/op
BenchmarkCache2Go/Get-4                  2000000               546 ns/op              39 B/op          3 allocs/op
BenchmarkGoCache/Set-4                   1000000              2059 ns/op             279 B/op          5 allocs/op
BenchmarkGoCache/Get-4                   3000000               449 ns/op              23 B/op          2 allocs/op
BenchmarkFreecache/Set-4                 1000000              1356 ns/op              80 B/op          4 allocs/op
BenchmarkFreecache/Get-4                 5000000               331 ns/op              24 B/op          2 allocs/op
BenchmarkBigCache/Set-4                  1000000              1029 ns/op              99 B/op          4 allocs/op
BenchmarkBigCache/Get-4                  3000000               426 ns/op              39 B/op          2 allocs/op
BenchmarkGCache/Set-4                    1000000              1227 ns/op             207 B/op          8 allocs/op
BenchmarkGCache/Get-4                    5000000               358 ns/op              39 B/op          3 allocs/op
BenchmarkSyncMap/Set-4                   1000000              1959 ns/op             258 B/op          9 allocs/op
BenchmarkSyncMap/Get-4                   5000000               299 ns/op              23 B/op          2 allocs/op


_Note: sync.Map does not support expiration and is only included for comparison_


## Warning

Please note that these libraries are benchmarked against storage of small strings. If you are storing large blobs the results _will_ differ and you are encouraged to benchmark with your custom payloads.

- https://golang.org/pkg/sync/#Map
- https://github.com/coocood/freecache
- https://github.com/allegro/bigcache
- https://github.com/patrickmn/go-cache
- https://github.com/muesli/cache2go
- https://github.com/bluele/gcache
