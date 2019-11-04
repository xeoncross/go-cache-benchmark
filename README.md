## Go Cache Benchmarks

Benchmarks of in-memory cache libraries for Go with expiration/TTL support.

## Run Test

    go get
    go test -bench=. -benchmem

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/Xeoncross/go-cache-benchmark
BenchmarkKodingCache/Set-8         	 1000000	      1357 ns/op	     424 B/op	       5 allocs/op
BenchmarkKodingCache/Get-8         	 2061524	       585 ns/op	      23 B/op	       2 allocs/op
BenchmarkHashicorpLRU/Set-8        	 2028216	       592 ns/op	     166 B/op	       8 allocs/op
BenchmarkHashicorpLRU/Get-8        	 5974563	       202 ns/op	      40 B/op	       3 allocs/op
BenchmarkCache2Go/Set-8            	 1000000	      1284 ns/op	     370 B/op	       9 allocs/op
BenchmarkCache2Go/Get-8            	 2805157	       392 ns/op	      39 B/op	       3 allocs/op
BenchmarkGoCache/Set-8             	 1000000	      1050 ns/op	     263 B/op	       5 allocs/op
BenchmarkGoCache/Get-8             	 4109612	       287 ns/op	      23 B/op	       2 allocs/op
BenchmarkFreecache/Set-8           	 1814077	       682 ns/op	      49 B/op	       4 allocs/op
BenchmarkFreecache/Get-8           	 4993617	       240 ns/op	      24 B/op	       2 allocs/op
BenchmarkBigCache/Set-8            	 1683318	       656 ns/op	      55 B/op	       4 allocs/op
BenchmarkBigCache/Get-8            	 2758429	       399 ns/op	      43 B/op	       3 allocs/op
BenchmarkGCache/Set-8              	 2181370	       597 ns/op	     191 B/op	       8 allocs/op
BenchmarkGCache/Get-8              	 5746874	       207 ns/op	      39 B/op	       3 allocs/op
BenchmarkRistretto/Set-8           	 1467294	       975 ns/op	     151 B/op	       7 allocs/op
BenchmarkRistretto/Get-8           	 2731618	       431 ns/op	      47 B/op	       3 allocs/op
BenchmarkSyncMap/Set-8             	 1000000	      1131 ns/op	     242 B/op	       9 allocs/op
BenchmarkSyncMap/Get-8             	 4409548	       267 ns/op	      23 B/op	       2 allocs/op
PASS
ok  	github.com/Xeoncross/go-cache-benchmark	45.898s
```

_Note: sync.Map does not support expiration and is only included for comparison_
_Note: hashicorp/golang-lru does not support expires/TTL and is only included for comparison_


## Warning

Please note that these libraries are benchmarked against storage of small strings. If you are storing large blobs the results _will_ differ and you are encouraged to benchmark with your custom payloads.

- https://github.com/dgraph-io/ristretto
- https://golang.org/pkg/sync/#Map
- https://github.com/coocood/freecache
- https://github.com/allegro/bigcache
- https://github.com/patrickmn/go-cache
- https://github.com/muesli/cache2go
- https://github.com/bluele/gcache
