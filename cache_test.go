package gocachebenchmarks

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/bluele/gcache"
	"github.com/coocood/freecache"
	hashicorp "github.com/hashicorp/golang-lru"
	koding "github.com/koding/cache"
	"github.com/muesli/cache2go"
	cache "github.com/patrickmn/go-cache"
	ristretto "github.com/dgraph-io/ristretto"
)

func toKey(i int) string {
	return fmt.Sprintf("item:%d", i)
}

// We will be storing many short strings as the key and value
func BenchmarkKodingCache(b *testing.B) {
	c := koding.NewMemoryWithTTL(time.Duration(60) * time.Second)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(toKey(i), toKey(i))
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			value, err := c.Get(toKey(i))
			if err == nil {
				_ = value
			}
		}
	})
}

func BenchmarkHashicorpLRU(b *testing.B) {
	// c := cache2go.Cache("test")
	c, _ := hashicorp.New(10)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Add(toKey(i), toKey(i))
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			value, err := c.Get(toKey(i))
			if err == true {
				_ = value
			}
		}
	})
}

func BenchmarkCache2Go(b *testing.B) {
	c := cache2go.Cache("test")

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Add(toKey(i), 1*time.Minute, toKey(i))
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Value(toKey(i))
			if err == nil {
				_ = value
			}
		}
	})
}

func BenchmarkGoCache(b *testing.B) {
	c := cache.New(1*time.Minute, 5*time.Minute)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Add(toKey(i), toKey(i), cache.DefaultExpiration)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, found := c.Get(toKey(i))
			if found {
				_ = value
			}
		}
	})
}

func BenchmarkFreecache(b *testing.B) {
	c := freecache.NewCache(1024 * 1024 * 5)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set([]byte(toKey(i)), []byte(toKey(i)), 60)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Get([]byte(toKey(i)))
			if err == nil {
				_ = value
			}
		}
	})
}

func BenchmarkBigCache(b *testing.B) {
	c, _ := bigcache.NewBigCache(bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,
		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 10,
	})

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(toKey(i), []byte(toKey(i)))
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Get(toKey(i))
			if err == nil {
				_ = value
			}
		}
	})
}

func BenchmarkGCache(b *testing.B) {
	c := gcache.New(b.N).LRU().Build()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(toKey(i), toKey(i))
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Get(toKey(i))
			if err == nil {
				_ = value
			}
		}
	})
}

func BenchmarkRistretto(b *testing.B) {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	
	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Set(toKey(i), toKey(i), 1)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, ok := cache.Get(toKey(i))
			if ok {
				_ = value
			}
		}
	})
	
}


// No expire, but helps us compare performance
func BenchmarkSyncMap(b *testing.B) {
	var m sync.Map

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(toKey(i), toKey(i))
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, found := m.Load(toKey(i))
			if found {
				_ = value
			}
		}
	})
}
