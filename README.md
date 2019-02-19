## Go Cache Benchmarks

Benchmarks of in-memory cache libraries for Go with expiration/TTL support.

## Run Test

    go get
    go test -bench=. -benchmem

## Results

    goos: darwin
    goarch: amd64
    pkg: github.com/Xeoncross/go-cache-benchmarks
    BenchmarkCache2Go/Set-8         	 1000000	      1622 ns/op	     370 B/op	       9 allocs/op
    BenchmarkCache2Go/Get-8         	 3000000	       386 ns/op	      39 B/op	       3 allocs/op
    BenchmarkGoCache/Set-8          	 1000000	      1061 ns/op	     279 B/op	       5 allocs/op
    BenchmarkGoCache/Get-8          	 5000000	       283 ns/op	      23 B/op	       2 allocs/op
    BenchmarkFreecache/Set-8        	 2000000	       775 ns/op	      66 B/op	       4 allocs/op
    BenchmarkFreecache/Get-8        	 5000000	       265 ns/op	      24 B/op	       2 allocs/op
    BenchmarkBigCache/Set-8         	 2000000	       786 ns/op	      99 B/op	       4 allocs/op
    BenchmarkBigCache/Get-8         	 3000000	       409 ns/op	      55 B/op	       3 allocs/op
    BenchmarkGCache/Set-8           	 2000000	       649 ns/op	     207 B/op	       8 allocs/op
    BenchmarkGCache/Get-8           	10000000	       228 ns/op	      39 B/op	       3 allocs/op
    BenchmarkSyncMap/Set-8          	 1000000	      1245 ns/op	     258 B/op	       9 allocs/op
    BenchmarkSyncMap/Get-8          	 5000000	       262 ns/op	      23 B/op	       2 allocs/op

_Note: sync.Map does not support expiration and is only included for comparison_


## Warning

Please note that these libraries are benchmarked against storage of small strings. If you are storing large blobs the results _will_ differ and you are encouraged to benchmark with your custom payloads.

- https://golang.org/pkg/sync/#Map
- https://github.com/coocood/freecache
- https://github.com/allegro/bigcache
- https://github.com/patrickmn/go-cache
- https://github.com/muesli/cache2go
- https://github.com/bluele/gcache
