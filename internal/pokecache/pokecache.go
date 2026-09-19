// Package pokecache
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	C  map[string]cacheEntry
	Mu sync.RWMutex
}

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{C: make(map[string]cacheEntry)}
	go cache.ReadLoop(interval)
	return cache
}

func (cache *Cache) Get(key string) (val []byte, found bool) {
	cache.Mu.RLock()
	defer cache.Mu.RUnlock()
	entry, found := cache.C[key]
	return entry.Val, found
}

func (cache *Cache) Add(key string, val []byte) {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()
	cache.C[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (cache *Cache) ReadLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		for key, value := range cache.C {
			res := time.Now().Compare(value.CreatedAt.Add(interval))
			if res == 0 || res == 1 {
				delete(cache.C, key)
			}
		}
	}
}
