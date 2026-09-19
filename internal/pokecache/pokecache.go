// Package pokecache
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	c  map[string]cacheEntry
	mu sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{c: make(map[string]cacheEntry)}
	go cache.ReadLoop(interval)
	return cache
}

func (cache *Cache) Get(key string) (val []byte, found bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	entry, found := cache.c[key]
	return entry.val, found
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.c[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) ReadLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.update(interval)
	}
}

func (cache *Cache) update(interval time.Duration) {
	cache.mu.Lock()
	for key, value := range cache.c {
		res := time.Now().Compare(value.createdAt.Add(interval))
		if res == 0 || res == 1 {
			delete(cache.c, key)
		}
	}
	cache.mu.Unlock()
}
