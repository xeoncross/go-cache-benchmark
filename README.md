## Go Cache Benchmarks

Benchmarks of in-memory cache libraries for Go with expiration/TTL support.

## Run Test

    go get
    go test -bench=. -benchmem

## Results

    goos: darwin
    goarch: amd64
    pkg: github.com/Xeoncross/go-cache-benchmarks
    BenchmarkKodingCache/Set-8         	 1000000	      1535 ns/op	     424 B/op	       5 allocs/op
    BenchmarkKodingCache/Get-8         	 3000000	       605 ns/op	      23 B/op	       2 allocs/op
    BenchmarkHashicorpLRU/Set-8        	 3000000	       581 ns/op	     166 B/op	       8 allocs/op
    BenchmarkHashicorpLRU/Get-8        	10000000	       202 ns/op	      40 B/op	       3 allocs/op
    BenchmarkCache2Go/Set-8            	 1000000	      1448 ns/op	     354 B/op	       9 allocs/op
    BenchmarkCache2Go/Get-8            	 3000000	       380 ns/op	      39 B/op	       3 allocs/op
    BenchmarkGoCache/Set-8             	 2000000	       923 ns/op	     196 B/op	       6 allocs/op
    BenchmarkGoCache/Get-8             	 5000000	       329 ns/op	      23 B/op	       2 allocs/op
    BenchmarkFreecache/Set-8           	 2000000	       651 ns/op	      49 B/op	       4 allocs/op
    BenchmarkFreecache/Get-8           	10000000	       231 ns/op	      24 B/op	       2 allocs/op
    BenchmarkBigCache/Set-8            	 2000000	       665 ns/op	      83 B/op	       4 allocs/op
    BenchmarkBigCache/Get-8            	 3000000	       375 ns/op	      45 B/op	       3 allocs/op
    BenchmarkGCache/Set-8              	 3000000	       640 ns/op	     191 B/op	       8 allocs/op
    BenchmarkGCache/Get-8              	10000000	       227 ns/op	      39 B/op	       3 allocs/op
    BenchmarkSyncMap/Set-8             	 1000000	      1249 ns/op	     242 B/op	       9 allocs/op
    BenchmarkSyncMap/Get-8             	 5000000	       263 ns/op	      23 B/op	       2 allocs/op


_Note: sync.Map does not support expiration and is only included for comparison_
_Note: hashicorp/golang-lru does not support expires/TTL and is only included for comparison_


## Warning

Please note that these libraries are benchmarked against storage of small strings. If you are storing large blobs the results _will_ differ and you are encouraged to benchmark with your custom payloads.

- https://golang.org/pkg/sync/#Map
- https://github.com/coocood/freecache
- https://github.com/allegro/bigcache
- https://github.com/patrickmn/go-cache
- https://github.com/muesli/cache2go
- https://github.com/bluele/gcache
